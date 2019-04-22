package cluster

type ClusterNodeInfo struct {
	NodeUuid string
	Meta map[string]string
	Tick int64
}

type INodeMgr interface {
	IsNodeActive (nodeUuid string) bool
	GetNodeInfo(nodeUuid string) *ClusterNodeInfo
	GetAllNodesInfo() []*ClusterNodeInfo
}

type NodeMgr struct {
	nodeMgr INodeMgr
}

func NewNodeMgr(nodeMgr INodeMgr) *NodeMgr {
	return &NodeMgr{nodeMgr: nodeMgr}
}

func (this *NodeMgr) IsNodeActive(nodeUuid string) bool {
	return this.nodeMgr.IsNodeActive(nodeUuid)
}

func (this *NodeMgr) GetNodeInfo(nodeUuid string) *ClusterNodeInfo {
	return this.nodeMgr.GetNodeInfo(nodeUuid)
}

func (this *NodeMgr) GetAllNodesInfo() []*ClusterNodeInfo {
	return this.nodeMgr.GetAllNodesInfo()
}
