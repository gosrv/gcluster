package route

import (
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gbase/gproto"
)

type routeMap struct {
	routes          map[interface{}][]gproto.FProcessor
	routeDefault    gproto.FProcessor
	single          bool
	showNoRouteWarn bool
}

func NewRouteMap(single bool, showNoRouteWarn bool) gproto.IRoute {
	return &routeMap{
		routes:          make(map[interface{}][]gproto.FProcessor),
		single:          single,
		showNoRouteWarn: showNoRouteWarn,
	}
}

func (this *routeMap) Connect(key interface{}, processor gproto.FProcessor) {
	if key == nil {
		this.routeDefault = processor
		return
	}
	if this.single && len(this.routes[key]) > 0 {
		glog.Panic("duplicate route key %v connect", key)
	}
	this.routes[key] = append(this.routes[key], processor)
}

func (this *routeMap) GetRoute(key interface{}) []gproto.FProcessor {
	return this.routes[key]
}

func (this *routeMap) Trigger(from interface{}, key interface{}, val interface{}) interface{} {
	route := this.routes[key]
	if len(route) > 0 {
		var rv interface{}
		for _, r := range route {
			r := r(from, key, val)
			if r != nil {
				rv = r
			}
		}
		return rv
	} else if this.routeDefault != nil {
		return this.routeDefault(from, key, val)
	} else if this.showNoRouteWarn {
		glog.Debug("no route for key %v", key)
	}
	return nil
}
