package gproto

type INetConnector interface {
	Connect(host string) INetChannel
	ConnectBind(localhost string, host string) INetChannel
}
