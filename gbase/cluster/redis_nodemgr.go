package cluster

import (
	"encoding/json"
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/gcluster/gbase/gutil"
	"time"
)

var clusterNodeKey = "_cluster._nodes"

type ActiveNodeInfo struct {
	nodeInfo *ClusterNodeInfo
	activeTime int64
}

type RedisNodeMgr struct {
	nodeUuid string					`bean:"app.id"`
	myNodeInfo *ClusterNodeInfo
	nodeOpt *gredis.HashOperation	`redis:""`
	activeNodesInfo map[string]*ActiveNodeInfo
	allNodesInfo []*ClusterNodeInfo
}

func (this *RedisNodeMgr) BeanInit() {
	this.myNodeInfo = &ClusterNodeInfo{NodeUuid:this.nodeUuid, Tick:time.Now().Unix()}
}

func (this *RedisNodeMgr) BeanUninit() {

}

func NewRedisNodeMgr() *RedisNodeMgr {
	return &RedisNodeMgr{}
}

func (this *RedisNodeMgr) GetNodeInfo(nodeUuid string) *ClusterNodeInfo {
	nodeInfo, ok := this.activeNodesInfo[nodeUuid]
	if !ok {
		return nil
	}
	return nodeInfo.nodeInfo
}

func (this *RedisNodeMgr) GetAllNodesInfo() []*ClusterNodeInfo {
	return this.allNodesInfo
}

func (this *RedisNodeMgr) IsNodeActive (nodeUuid string) bool {
	_, ok := this.activeNodesInfo[nodeUuid]
	return ok
}

func (this *RedisNodeMgr) GOStart() {
	gutil.RecoverGo(func() {
		for {
			this.registerServer()
			time.Sleep(time.Second * 2)
		}
	})
}

func (this *RedisNodeMgr) registerServer() {
	now := time.Now().Unix()
	this.myNodeInfo.Tick = now
	data, err := json.Marshal(this.myNodeInfo)
	if err != nil {
		gl.Debug("json marshal error %v", err)
	}
	_, err = this.nodeOpt.HSet(clusterNodeKey, this.myNodeInfo.NodeUuid, string(data))
	if err != nil {
		gl.Debug("redis hash opt error %v", err)
	}
	allData, err := this.nodeOpt.HGetAll(clusterNodeKey)
	if err != nil {
		gl.Debug("redis hash opt error %v", err)
	}
	nodesInfo := make(map[string]*ActiveNodeInfo)
	for uuid, ndata := range allData {
		activeNodeInfo := &ActiveNodeInfo{}
		err := json.Unmarshal([]byte(ndata), &activeNodeInfo.nodeInfo)
		if err != nil {
			gl.Debug("json unmarshal error %v", err)
			continue
		}
		oldActiveNodeInfo, ok := this.activeNodesInfo[uuid]
		if !ok || oldActiveNodeInfo.nodeInfo.Tick != activeNodeInfo.nodeInfo.Tick {
			activeNodeInfo.activeTime = now
			nodesInfo[uuid] = activeNodeInfo
		} else if oldActiveNodeInfo.activeTime + 10 < now {
			this.nodeOpt.HDel(clusterNodeKey, uuid)
		} else {
			nodesInfo[uuid] = oldActiveNodeInfo
		}
	}
	allNodesInfo := make([]*ClusterNodeInfo, 0, len(nodesInfo))
	for _, ni := range nodesInfo {
		allNodesInfo = append(allNodesInfo, ni.nodeInfo)
	}
	this.activeNodesInfo = nodesInfo
	this.allNodesInfo = allNodesInfo
}