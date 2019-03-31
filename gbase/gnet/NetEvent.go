package gnet

const (
	NetEventDisconnect = -1 //断开连接
	NetEventConnect    = -2 //建立连接

	NetEventReadException  = -3  // 读取异常，输入异常
	NetEventWriteException = -4  // 写入异常，输出异常
	NetEventReadIdle       = -5  // 读取心跳超时
	NetEventWriteIdle      = -6  // 写入心跳超时
	NetEventTick           = -11 // 心跳
)
