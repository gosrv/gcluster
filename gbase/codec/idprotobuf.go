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
	HeadLenIdProtobuf = 2
)

// 网络消息pb编码器，用于websocket数据编码
// 2字节proto id + proto
type IdProtobufEncoder struct {
	// id 类型指针映射
	idtype gnet.ITypeID
}

func NewIdProtobufEncoder(idtype gnet.ITypeID) gproto.IEncoder {
	return &IdProtobufEncoder{
		idtype: idtype,
	}
}

func (this *IdProtobufEncoder) Encode(val interface{}) interface{} {
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

	totalLen := HeadLenIdProtobuf + len(pbData)
	totalData := make([]byte, totalLen)
	_, err = gutil.IntCodec.PackFixInt16BE(int16(id), totalData, 0)
	util.VerifyNoError(err)

	copy(totalData[HeadLenIdProtobuf:], pbData)
	return totalData
}

type IdProtobufDecoder struct {
	idtype gnet.ITypeID
}

// 网络消息pb解码器
// 2字节proto id + proto
func NewIdProtobufDecoder(idtype gnet.ITypeID) gproto.IDecoder {
	return &IdProtobufDecoder{
		idtype: idtype,
	}
}

func (this *IdProtobufDecoder) Decode(input interface{}) interface{} {
	data := input.([]byte)
	util.Verify(len(data) >= HeadLenIdProtobuf)
	id, err := gutil.IntCodec.UnpackFixInt16BE(data, 2)
	util.VerifyNoError(err)

	tp, err := this.idtype.ID2Type(int(id))
	if err != nil {
		gl.Panic("can not find id type %v:%v", id, err)
		return nil
	}

	value := reflect.New(tp.Elem()).Interface().(proto.Message)
	err = proto.Unmarshal(data[HeadLenIdProtobuf:], value)
	if err != nil {
		gl.Panic("proto unmarshal error %v:%v", tp, hex.EncodeToString(data[HeadLenIdProtobuf:]))
	}
	return value
}
