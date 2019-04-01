package entity

import (
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/goioc/util"
	"reflect"
)

/**
网络消息类型和id的映射关系建立
*/
func NewLogicMsgIds() gnet.ITypeID {
	msgIds := gnet.NewTypeID()

	err := msgIds.AddIDType(1, reflect.TypeOf((*netproto.PlayerData)(nil)))
	util.VerifyNoError(err)
	err = msgIds.AddIDType(2, reflect.TypeOf((*netproto.PlayerInfo)(nil)))
	util.VerifyNoError(err)

	for id, name := range netproto.EMsgIds_name {
		err := msgIds.AddIDType(int(id), proto.MessageType("netproto."+name[1:]))
		util.VerifyNoError(err)
	}

	return msgIds
}
