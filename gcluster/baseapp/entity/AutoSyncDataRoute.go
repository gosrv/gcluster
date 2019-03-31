package entity

import (
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gbase/gutil"
	"reflect"
)

type AutoSyncDataRoute struct {
	raw        gproto.IRoute
	playerSync *PlayerDataSync
}

func NewAutoSyncDataRoute() *AutoSyncDataRoute {
	return &AutoSyncDataRoute{}
}

func (this *AutoSyncDataRoute) SetDelegate(raw gproto.IRoute) {
	this.raw = raw
}

func (this *AutoSyncDataRoute) Connect(key interface{}, processor gproto.FProcessor) {
	glog.Panic("not support operation")
}

func (this *AutoSyncDataRoute) GetRoute(key interface{}) []gproto.FProcessor {
	return this.raw.GetRoute(key)
}

func (this *AutoSyncDataRoute) Trigger(from interface{}, key interface{}, value interface{}) interface{} {
	ctx := from.(gnet.ISessionCtx)
	rep := this.raw.Trigger(from, key, value)
	if this.playerSync == nil {
		psync := ctx.Get(reflect.TypeOf((*PlayerDataSync)(nil)))
		if psync != nil {
			this.playerSync = psync.(*PlayerDataSync)
		}
	}
	if this.playerSync != nil {
		this.playerSync.TrySyncDirtyData(false)
	}
	if gutil.IsNilValue(rep) {
		glog.Debug("msg process %v:%v	->	nil", key, value)
	} else {
		glog.Debug("msg process %v:%v	->	%v:%v",
			key, value, reflect.TypeOf(rep), rep)
	}
	return rep
}
