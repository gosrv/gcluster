package service

import (
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/tcpnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"reflect"
)

type serviceDataAutoSync struct {
	net *tcpnet.TcpNetServer `bean:""`
}

func NewServiceDataAutoSync() *serviceDataAutoSync {
	return &serviceDataAutoSync{}
}

func (this *serviceDataAutoSync) BeanInit() {
	eventRoute := this.net.GetEventRoute()
	eventRoute.Connect(gnet.NetEventTick, func(from interface{}, key interface{}, value interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		sync := ctx.Get(reflect.TypeOf((*entity.PlayerDataSync)(nil)))
		if sync != nil {
			sync.(*entity.PlayerDataSync).TrySyncDirtyData(false)
		}

		return nil
	})
	eventRoute.Connect(gnet.NetEventDisconnect, func(from interface{}, key interface{}, value interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		sync := ctx.Get(reflect.TypeOf((*entity.PlayerDataSync)(nil))).(*entity.PlayerDataSync)
		sync.TrySyncDirtyData(true)
		return nil
	})
}

func (this *serviceDataAutoSync) BeanUninit() {

}
