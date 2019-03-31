package gnet

import (
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gbase/route"
)

type NetConfig struct {
	// 网络消息编码器
	Encoder gproto.IEncoder
	// 网络消息解码器
	Decoder gproto.IDecoder
	// 事件路由器
	EventRoute gproto.IRoute
	// 消息数据路由器
	DataRoute gproto.IRoute
	// 读缓冲区大小
	ReadBufSize int
	// 写消息队列大小
	WriteChannelSize int
	HeartTickMs      int
}

func NewNetConfig() *NetConfig {
	config := &NetConfig{
		EventRoute:       route.NewRouteMap(false, false),
		DataRoute:        route.NewRouteMap(true, true),
		ReadBufSize:      16384,
		WriteChannelSize: 1024,
		HeartTickMs:      10000,
	}

	return config
}
