package common

type ClusterNodeMgr struct {
}

func NewClusterNodeMgr() *ClusterNodeMgr {
	return &ClusterNodeMgr{}
}

func (this *ClusterNodeMgr) IsNodeActive(nodeUuid string) bool {
	return false
}

func (this *ClusterNodeMgr) SendMsgToNode(nodeUuid string) {

}
