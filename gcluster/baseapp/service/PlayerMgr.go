package service

import (
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gcluster/gbase/gdb"
	"github.com/gosrv/gcluster/gbase/gdb/dbaccessor"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/common/meta"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/goioc/util"
	"github.com/sirupsen/logrus"
	"reflect"
	"sort"
	"strconv"
	"sync"
)

type PlayerMgr struct {
	playerId2Channel sync.Map
	log              *logrus.Logger `log:"app"`

	messageQueueFactories   []gdb.IMessageQueueFactory     `bean:""`
	attributeGroupFactories []gdb.IDBAttributeGroupFactory `bean:""`
	dbAccessorFactory       *dbaccessor.DBDataAccessorFactory
}

func (this *PlayerMgr) BeanInit() {
	if len(this.messageQueueFactories) == 0 {
		util.Panic("no message queue found")
	}
	if len(this.attributeGroupFactories) == 0 {
		util.Panic("no attribute group found")
	}
	sort.Slice(this.messageQueueFactories, func(i, j int) bool {
		return this.messageQueueFactories[i].GetPriority() < this.messageQueueFactories[j].GetPriority()
	})
	sort.Slice(this.attributeGroupFactories, func(i, j int) bool {
		return this.attributeGroupFactories[i].GetPriority() < this.attributeGroupFactories[j].GetPriority()
	})
	this.dbAccessorFactory = dbaccessor.NewDBDataAccessorFactory(this.messageQueueFactories[0], this.attributeGroupFactories)
}

func (this *PlayerMgr) BeanUninit() {

}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{}
}

func (this *PlayerMgr) IsForbidLogin(playerId int64) bool {
	return false
}

func (this *PlayerMgr) GetDBDataAccessor(playerId int64) *dbaccessor.DBDataAccessor {
	return this.dbAccessorFactory.GetDataAccessor(meta.PlayerAttribute, strconv.FormatInt(playerId, 10))
}

func (this *PlayerMgr) LoadPlayerData(playerId int64, loader *gdb.TheDataLoaderChain) *entity.PlayerData {
	val, err := loader.Load()
	if err != nil {
		this.log.Debugf("load player data error %v", err)
	}

	pd := entity.NewPlayerData()
	if len(val) == 0 {
		this.InitPlayerData(playerId, pd)
	} else {
		npd := &netproto.PlayerData{}
		err = proto.Unmarshal([]byte(val), npd)
		if err != nil {
			this.log.Debugf("unmarshal player data error %v", err)
		}
		pd.FromProto(npd)
	}
	return pd
}

func (this *PlayerMgr) InitPlayerData(playerId int64, playerData *entity.PlayerData) {
	baseInfo := playerData.GetBaseInfo()
	baseInfo.SetId(playerId)
	baseInfo.SetLevel(1)
	baseInfo.SetName("name" + strconv.FormatInt(playerId, 10))
}

func (this *PlayerMgr) PlayerLogin(playerId int64, netChannel gproto.INetChannel, ctx gnet.ISessionCtx, dataAccessor *dbaccessor.DBDataAccessor) bool {
	_, loaded := this.playerId2Channel.LoadOrStore(playerId, netChannel)
	if loaded {
		return false
	}
	attributeGroup := dataAccessor.GetAttributeGroup()
	loader := dataAccessor.GetDataLoader(meta.PlayerBaseData)
	saver := dataAccessor.GetDataSaver(meta.PlayerBaseData)
	playerData := this.LoadPlayerData(playerId, loader)
	playerInfo := entity.NewPlayerInfo()
	onlineData := entity.NewPlayerOnlineData()
	syncData := entity.NewPlayerDataSync(playerData, playerInfo,
		20, 60, netChannel,
		saver, attributeGroup)

	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(playerData), playerData)
	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(playerInfo), playerInfo)
	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(onlineData), onlineData)
	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(syncData), syncData)

	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(dataAccessor), dataAccessor)
	ctx.SetAttribute(gnet.ScopeSession, gdb.IDBAttributeGroupType, attributeGroup)
	ctx.SetAttribute(gnet.ScopeSession, gdb.IMessageQueueType,
		this.dbAccessorFactory.GetMessageQueue(meta.PlayerAttribute, strconv.FormatInt(playerId, 10)))

	return true
}

func (this *PlayerMgr) PlayerLogout(playerId int64) {
	this.playerId2Channel.Delete(playerId)
}

func (this *PlayerMgr) GetNetchannelByPlayerId(playerId int64) gproto.INetChannel {
	net, ok := this.playerId2Channel.Load(playerId)
	if !ok {
		return nil
	}
	return net.(gproto.INetChannel)
}
