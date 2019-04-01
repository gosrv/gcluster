package route

import (
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gbase/gproto"
)

type RouteDelegate struct {
	raw       gproto.IRoute
	modifible bool
}

func NewRouteDelegate(modifible bool) *RouteDelegate {
	return &RouteDelegate{modifible: modifible}
}

func (this *RouteDelegate) GetRouteKeys() []interface{} {
	return this.raw.GetRouteKeys()
}

func (this *RouteDelegate) SetDelegate(raw gproto.IRoute) {
	this.raw = raw
}

func (this *RouteDelegate) Connect(key interface{}, processor gproto.FProcessor) {
	if !this.modifible {
		glog.Panic("not support operation")
	} else {
		this.raw.Connect(key, processor)
	}
}

func (this *RouteDelegate) GetRoute(key interface{}) []gproto.FProcessor {
	return this.raw.GetRoute(key)
}

func (this *RouteDelegate) Trigger(from interface{}, key interface{}, value interface{}) interface{} {
	return this.raw.Trigger(from, key, value)
}
