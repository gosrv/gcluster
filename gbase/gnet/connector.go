package gnet

type INetConnector interface {
	NetConnect(address string)
}
