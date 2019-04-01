package common

import (
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gcluster/common/meta"
	"strconv"
	"time"
)

type ClusterMsgCenter struct {
	msgListOpt *gredis.ListOperation `redis.obj:""`
	encoder    gproto.IEncoder
	decoder    gproto.IDecoder
}

func NewClusterMsgCenter(encoder gproto.IEncoder, decoder gproto.IDecoder) *ClusterMsgCenter {
	return &ClusterMsgCenter{encoder: encoder, decoder: decoder}
}

func (this *ClusterMsgCenter) getPlayerMsgQueueName(playerId int64) string {
	return meta.MessageQueuePlayer + ":" + strconv.FormatInt(playerId, 10)
}

func (this *ClusterMsgCenter) getBundleMsgQueueName(bundleId int64) string {
	return meta.MessageQueueBundle + ":" + strconv.FormatInt(bundleId, 10)
}

func (this *ClusterMsgCenter) BlockProcessPlayerMsg(id int64, processor gproto.IMsgProcessor) {
	queue := this.getPlayerMsgQueueName(id)
	this.blockProcessMsg(queue, processor)
}

func (this *ClusterMsgCenter) BlockProcessBundleMsg(id int64, processor gproto.IMsgProcessor) {
	queue := this.getBundleMsgQueueName(id)
	this.blockProcessMsg(queue, processor)
}

func (this *ClusterMsgCenter) blockProcessMsg(mqName string, processor gproto.IMsgProcessor) {
	for {
		vals, _ := this.msgListOpt.BLeftPop(time.Minute, mqName)
		for _, val := range vals {
			msg := this.decoder.Decode(val)
			processed := processor.ProcessMsg(msg)
			if !processed {
				this.msgListOpt.LeftPush(mqName, val)
				return
			}
		}
	}
}

func (this *ClusterMsgCenter) SendToPlayer(playerId int64, msg interface{}) error {
	return this.sendToTarget(this.getPlayerMsgQueueName(playerId), msg)
}

func (this *ClusterMsgCenter) SendToBundle(bundleId int64, msg interface{}) error {
	return this.sendToTarget(this.getBundleMsgQueueName(bundleId), msg)
}

func (this *ClusterMsgCenter) sendToTarget(target string, msg interface{}) error {
	_, err := this.msgListOpt.RightPush(target, this.encoder.Encode(msg))
	return err
}
