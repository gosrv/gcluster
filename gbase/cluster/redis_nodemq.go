package cluster

import (
	"fmt"
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gbase/gutil"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

type RedisNodeMQ struct {
	listOperation *gredis.ListOperation	`redis:""`
	clusterNodeMgr INodeMgr
	encoder gproto.IEncoder
	decoder gproto.IDecoder
	route gproto.IRoute
	selfNodeUuid string	`bean:"app.id"`
}

func NewRedisNodeMQ(clusterNodeMgr INodeMgr, encoder gproto.IEncoder, decoder gproto.IDecoder, route gproto.IRoute) *RedisNodeMQ {
	return &RedisNodeMQ{clusterNodeMgr: clusterNodeMgr, encoder: encoder, decoder: decoder, route: route}
}

func (this *RedisNodeMQ) GetNodeMQKey(nodeUuid string) string {
	return NodeMQName + ":" + nodeUuid
}

func (this *RedisNodeMQ) GOStart() {
	ctx := gnet.NewSessionCtx()
	gutil.RecoverGo(func() {
		for {
			msg, err := this.BPop(this.selfNodeUuid)
			if err != nil {
				gl.Debug("redis node mq error %v", err)
				continue
			}
			ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(msg), msg)
			this.route.Trigger(ctx, reflect.TypeOf(msg), msg)
		}
	})
}

func (this *RedisNodeMQ) BPop(nodeUuid string) (interface{}, error) {
	queueKey := this.GetNodeMQKey(nodeUuid)
	data, err := this.listOperation.BLPop(time.Minute, queueKey)
	if err != nil {
		return nil, err
	}
	return this.decoder.Decode(data), err
}

func (this *RedisNodeMQ) Push(nodeUuid string, msg interface{}) error {
	if len(nodeUuid) == 0 {
		return errors.New("node uuid cannt empty")
	}
	if !this.clusterNodeMgr.IsNodeActive(nodeUuid) {
		return fmt.Errorf("node %v is inactive", nodeUuid)
	}
	queueKey := this.GetNodeMQKey(nodeUuid)
	msgData := this.encoder.Encode(msg)
	_, err := this.listOperation.RPush(queueKey, msgData)
	return err
}

