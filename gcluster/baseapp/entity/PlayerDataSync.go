package entity

import (
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gcluster/gbase/datasync"
	"github.com/gosrv/gcluster/gbase/gdb"
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gbase/gproto"
	"reflect"
	"time"
)

const (
	IdxPlayerDataMonitor = 0
	IdxPlayerInfoMonitor = 1
)

type PlayerDataSync struct {
	// 玩家数据，会同步到客户端，会存入数据库
	data *PlayerData
	// 玩家数据，会同步到客户端，不会存入数据库
	info *PlayerInfo

	redisFlushGapSec int64
	mongoFlushGapSec int64

	netChannel gproto.INetChannel
	dataSaver  *gdb.TheDataSaverChain

	prefreshDailyDataTime time.Time
	dirtyMonitor          datasync.IDirtyContainerMark

	preFlushRedisTime int64
	preFlushMongoTime int64

	dataExpireable gdb.IDataExpireable
}

func NewPlayerDataSync(data *PlayerData, info *PlayerInfo, redisFlushGapSec int64,
	mongoFlushGapSec int64, netChannel gproto.INetChannel,
	dataSaver *gdb.TheDataSaverChain, attributeGroup gdb.IDBAttributeGroup) *PlayerDataSync {
	datasync := &PlayerDataSync{
		data:             data,
		info:             info,
		redisFlushGapSec: redisFlushGapSec,
		mongoFlushGapSec: mongoFlushGapSec,
		netChannel:       netChannel,
		dataSaver:        dataSaver,
		dirtyMonitor:     datasync.NewDirtyContainerVector(),
	}
	// 如果有多级属性存储，我们可以把第一级设为可过期
	// 这里使用了redis和mongo，相当于把redis设为热数据
	if reflect.TypeOf(attributeGroup).AssignableTo(gdb.IDataExpireableType) {
		datasync.dataExpireable = attributeGroup.(gdb.IDataExpireable)
	}
	datasync.data.Init(datasync.dirtyMonitor, IdxPlayerDataMonitor)
	datasync.info.Init(datasync.dirtyMonitor, IdxPlayerInfoMonitor)
	return datasync
}

func (this *PlayerDataSync) markDBDataDirty() {
	if this.preFlushRedisTime == 0 {
		this.preFlushRedisTime = time.Now().Unix()
	}
	if this.preFlushMongoTime == 0 {
		this.preFlushMongoTime = time.Now().Unix()
	}
}

func (this *PlayerDataSync) clearDBDataDirty() {
	this.preFlushRedisTime = 0
	this.preFlushMongoTime = 0
}

func (this *PlayerDataSync) TrySyncDirtyData(flushdb bool) {
	if this.dirtyMonitor.IsDirty(IdxPlayerInfoMonitor) {
		this.netChannel.Send(this.info.ToProtoDirty())
	}

	if this.dirtyMonitor.IsDirty(IdxPlayerDataMonitor) {
		this.markDBDataDirty()
		this.netChannel.Send(this.data.ToProtoDirty())
	}
	this.dirtyMonitor.ClearDirty()

	if this.preFlushMongoTime > 0 && time.Now().Unix()-this.preFlushRedisTime > this.redisFlushGapSec {
		this.clearDBDataDirty()
		bindata, err := proto.Marshal(this.data.ToProto())
		if err != nil {
			glog.Panic("save player data error %v", err)
		}
		this.dataSaver.SaveDepth(string(bindata), 0)
		if this.dataExpireable != nil {
			err = this.dataExpireable.SetExpireDuration(7 * 24 * time.Hour)
			if err != nil {
				glog.Warn("set cache data expire duration failed")
			}
		}
	}

	if this.preFlushRedisTime > 0 && time.Now().Unix()-this.preFlushRedisTime > this.redisFlushGapSec {
		this.preFlushRedisTime = 0
		bindata, err := proto.Marshal(this.data.ToProto())
		if err != nil {
			glog.Panic("save player data error %v", err)
		}
		this.dataSaver.SaveDepth(string(bindata), 1)
	}
}
