package tcpnet

import (
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"net"
	"sync"
	"time"
)

type netChannel struct {
	encoder gproto.IEncoder
	conn    net.Conn
	closed  bool
	// 未处理消息，需要通过路由器进行处理
	msgProcessedChannel chan interface{}
	// 已经处理过的消息，需要通过网络发出去
	msgUnprocessedChannel chan interface{}
	heartTicker           *time.Ticker
	closeChannel          chan int
	ctx                   gnet.ISessionCtx
	lock                  sync.Mutex
}

var _ gproto.INetChannel = (*netChannel)(nil)

func NewNetChannel(conn net.Conn,
	msgProcessedChannelSize int,
	msgUnprocessedChannelSize int,
	heartTickMs int,
) *netChannel {
	return &netChannel{
		conn:                  conn,
		msgProcessedChannel:   make(chan interface{}, msgProcessedChannelSize),
		msgUnprocessedChannel: make(chan interface{}, msgUnprocessedChannelSize),
		closed:                false,
		ctx:                   gnet.NewSessionCtx(),
		heartTicker:           time.NewTicker(time.Duration(heartTickMs) * time.Millisecond),
		closeChannel:          make(chan int, 1),
	}
}

func (this *netChannel) IsActive() bool {
	return !this.closed
}

func (this *netChannel) RemoteAddr() net.Addr {
	return this.conn.RemoteAddr()
}

func (this *netChannel) LocalAddr() net.Addr {
	return this.conn.LocalAddr()
}

func (this *netChannel) ProcessMsg(msg interface{}) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	this.msgUnprocessedChannel <- msg
	return false
}

func (this *netChannel) Send(data interface{}) {
	if this.encoder != nil {
		this.msgProcessedChannel <- this.encoder.Encode(data)
	} else {
		this.msgProcessedChannel <- data
	}
}

func (this *netChannel) Close() error {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.closed {
		return nil
	}
	this.closed = true
	this.heartTicker.Stop()
	close(this.closeChannel)
	close(this.msgUnprocessedChannel)
	return this.conn.Close()
}

func (this *netChannel) SetEncoder(encoder gproto.IEncoder) {
	this.encoder = encoder
}
