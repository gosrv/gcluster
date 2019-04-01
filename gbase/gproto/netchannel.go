package gproto

import (
	"net"
	"reflect"
)

type INetChannel interface {
	IMsgProcessor
	// 设置编码器
	SetEncoder(encoder IEncoder)
	// 发送一个数据
	Send(data interface{})
	// 关闭
	Close() error
	IsActive() bool
	RemoteAddr() net.Addr
	LocalAddr() net.Addr
}

var INetChannelType = reflect.TypeOf((*INetChannel)(nil)).Elem()
