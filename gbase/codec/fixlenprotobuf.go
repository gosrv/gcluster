package codec

import (
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gbase/gutil"
	"github.com/gosrv/goioc/util"
	"reflect"
)

const (
	HeadLenFixLenProtobuf = 6
)

// 网络消息固定长度pb编码器，用于tcp数据流编码
// 4字节长度 + 2字节proto id + proto
type NetMsgFixLenProtobufEncoder struct {
	// id 类型指针映射
	idtype gnet.ITypeID
}

func NewNetMsgFixLenProtobufEncoder(idtype gnet.ITypeID) gproto.IEncoder {
	return &NetMsgFixLenProtobufEncoder{
		idtype: idtype,
	}
}

func (this *NetMsgFixLenProtobufEncoder) Encode(val interface{}) interface{} {
	tp := reflect.TypeOf(val)
	id, err := this.idtype.Type2ID(tp)
	if err != nil {
		gl.Panic("can not get type id %v:%v", tp, err)
		return nil
	}

	pbData, err := proto.Marshal(val.(proto.Message))
	if err != nil {
		gl.Panic("proto marshal error %v:%v", tp, err)
		return nil
	}

	totalLen := HeadLenFixLenProtobuf + len(pbData)
	totalData := make([]byte, totalLen)
	_, err = gutil.IntCodec.PackFixInt32BE(int32(totalLen), totalData, 0)
	util.VerifyNoError(err)

	_, err = gutil.IntCodec.PackFixInt16BE(int16(id), totalData, 4)
	util.VerifyNoError(err)

	copy(totalData[HeadLenFixLenProtobuf:], pbData)
	return totalData
}

type NetMsgFixLenProtobufDecoder struct {
	idtype gnet.ITypeID
}

// 网络消息固定长度pb解码器
// 4字节长度 + 2字节proto id + proto
func NewNetMsgFixLenProtobufDecoder(idtype gnet.ITypeID) gproto.IDecoder {
	return &NetMsgFixLenProtobufDecoder{
		idtype: idtype,
	}
}

func (this *NetMsgFixLenProtobufDecoder) Decode(input interface{}) interface{} {
	reader := input.(*gutil.Buffer)
	if reader.Len() < HeadLenFixLenProtobuf {
		return nil
	}
	totalLen, err := gutil.IntCodec.UnpackFixInt32BE(reader.Peek(6), 0)
	util.VerifyNoError(err)

	if reader.Len() < int(totalLen) {
		return nil
	}
	totalData := reader.Read(int(totalLen))
	id, err := gutil.IntCodec.UnpackFixInt16BE(totalData, 4)
	util.VerifyNoError(err)

	tp, err := this.idtype.ID2Type(int(id))
	if err != nil {
		gl.Panic("can not find id type %v:%v", id, err)
		return nil
	}

	value := reflect.New(tp.Elem()).Interface().(proto.Message)
	err = proto.Unmarshal(totalData[HeadLenFixLenProtobuf:], value)
	if err != nil {
		gl.Panic("proto unmarshal error %v:%v", tp, hex.EncodeToString(totalData[HeadLenFixLenProtobuf:]))
	}
	return value
}
