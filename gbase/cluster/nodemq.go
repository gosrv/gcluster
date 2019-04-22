package cluster

var NodeMQName = "_node.mq"

type INodeMQ interface {
	Push(nodeUuid string, msg interface{}) error
}

type NodeMQ struct {
	nodeMq INodeMQ
}

func NewNodeMQ(nodeMq INodeMQ) *NodeMQ {
	return &NodeMQ{nodeMq: nodeMq}
}

func (this *NodeMQ) Push(nodeUuid string, msg interface{}) error {
	return this.nodeMq.Push(nodeUuid, msg)
}

