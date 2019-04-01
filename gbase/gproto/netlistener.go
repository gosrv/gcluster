package gproto

type INetListener interface {
	SetEncoder(encoder IEncoder)
	SetDecoder(decoder IDecoder)
	SetRoute(route IRoute)
	Listen(host string) error
}
