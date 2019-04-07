package tcpnet

import (
	"bufio"
	"errors"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gbase/gutil"
	"github.com/gosrv/goioc/util"
	"io"
	"net"
	"reflect"
	"sync/atomic"
)

const (
	CACHE_WRITE_BUFF_SIZE = 8192
)

func readProcess(netChannel *netChannel, config *gnet.NetConfig) error {
	decoder := config.Decoder
	buf := gutil.NewBuffer(config.ReadBufSize)
	for netChannel.IsActive() {
		_, err := buf.Fill(netChannel.conn)
		// 解析并处理所有的数据包
		for {
			netMsg := decoder.Decode(buf)
			if netMsg == nil {
				break
			}
			// 交由写线程处理
			netChannel.msgUnprocessedChannel <- netMsg
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func writeMessage(writer io.Writer, msg interface{}, config *gnet.NetConfig) error {
	msgbuffer := config.Encoder.Encode(msg)
	n, err := writer.Write(msgbuffer.([]byte))
	if err != nil {
		return err
	}
	if n != len(msgbuffer.([]byte)) {
		return errors.New("write not finished.")
	}
	return nil
}

func flushWriteMessage(writer *bufio.Writer, outChannelRead <-chan interface{}, config *gnet.NetConfig) error {
	encoder := config.Encoder
	for {
		select {
		case msg := <-outChannelRead:
			msgbuffer := encoder.Encode(msg)
			writer.Write(msgbuffer.([]byte))
		default:
			return nil
		}
	}
}

func writeProcess(netChannel *netChannel, config *gnet.NetConfig) error {
	eventRoute := config.EventRoute
	netDataWriter := bufio.NewWriterSize(netChannel.conn, CACHE_WRITE_BUFF_SIZE)
	dataRoute := config.DataRoute

	eventRoute.Trigger(netChannel.ctx, gnet.NetEventConnect, nil)
	defer func() {
		netChannel.Close()
		for {
			// 把剩下的活干完就可以退出了
			if msg, ok := <-netChannel.msgUnprocessedChannel; ok {
				netChannel.ctx.Clear(gnet.ScopeRequest)
				netChannel.ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(msg), msg)
				_ = dataRoute.Trigger(netChannel.ctx, reflect.TypeOf(msg), msg)
			} else {
				break
			}
		}
		eventRoute.Trigger(netChannel.ctx, gnet.NetEventDisconnect, nil)
	}()

	for netChannel.IsActive() {
		// 读操作可能被阻塞非常久，但是写操作不应该被阻塞太久
		// 如果写阻塞了，说明网络拥塞或客户端处理不过来了，也没必要处理其它数据了
		select {
		case _, _ = <-netChannel.closeChannel:
			return nil
		case msg, ok := <-netChannel.msgProcessedChannel:
			{
				if !ok {
					// channel已被关闭
					return errors.New("gnet msg processed channel closed.")
				}
				err := writeMessage(netDataWriter, msg, config)
				if err != nil {
					return err
				}
				err = flushWriteMessage(netDataWriter, netChannel.msgProcessedChannel, config)
				flusherr := netDataWriter.Flush()
				if err != nil {
					return err
				}
				if flusherr != nil {
					return flusherr
				}
			}
		case msg, ok := <-netChannel.msgUnprocessedChannel:
			{
				if !ok {
					// channel已被关闭
					return errors.New("gnet msg unprocessed channel closed.")
				}
				netChannel.ctx.Clear(gnet.ScopeRequest)
				netChannel.ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(msg), msg)
				response := dataRoute.Trigger(netChannel.ctx, reflect.TypeOf(msg), msg)
				if response == nil || reflect.ValueOf(response).IsNil() {
					continue
				}
				netChannel.msgProcessedChannel <- response
			}
		case _, ok := <-netChannel.heartTicker.C:
			{
				if !ok {
					return errors.New("ticker stoped")
				}
				eventRoute.Trigger(netChannel.ctx, gnet.NetEventTick, nil)
			}
		}
	}
	return nil
}

func GoListen(network string, address string, config *gnet.NetConfig) (net.Addr, error) {
	ln, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	gutil.RecoverGo(func() {
		for {
			conn, err := ln.Accept()
			util.VerifyNoError(err)
			netChannel := NewNetChannel(conn, config.WriteChannelSize,
				config.WriteChannelSize, config.HeartTickMs)
			netChannel.ctx.SetAttribute(gnet.ScopeSession, gproto.INetChannelType, netChannel)
			netChannel.ctx.SetAttribute(gnet.ScopeSession, gnet.ISessionCtxType, netChannel.ctx)
			// 每条连接由两个协程处理，一个读，一个写
			// reader
			gutil.RecoverGo(func() {
				defer netChannel.Close()

				_ = readProcess(netChannel, config)
			})

			// writer
			gutil.RecoverGo(func() {
				defer netChannel.Close()

				_ = writeProcess(netChannel, config)
			})
		}
	})

	return ln.Addr(), nil
}

func GoConnect(network, address string, config *gnet.NetConfig) (runover *int32) {
	ro := int32(0)
	runover = &ro

	atomic.StoreInt32(runover, 0)

	gutil.RecoverGo(func() {
		defer atomic.StoreInt32(runover, 1)

		conn, err := net.Dial(network, address)
		util.VerifyNoError(err)
		netChannel := NewNetChannel(conn, config.WriteChannelSize,
			config.WriteChannelSize, config.HeartTickMs)
		netChannel.ctx.SetAttribute(gnet.ScopeSession, gproto.INetChannelType, netChannel)
		netChannel.ctx.SetAttribute(gnet.ScopeSession, gnet.ISessionCtxType, netChannel.ctx)
		// reader
		gutil.RecoverGo(func() {
			defer func() {
				netChannel.Close()
			}()

			_ = readProcess(netChannel, config)
		})

		_ = writeProcess(netChannel, config)
	})

	return runover
}
