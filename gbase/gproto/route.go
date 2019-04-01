package gproto

import "reflect"

type FProcessor = func(interface{}, interface{}, interface{}) interface{}

type IRoute interface {
	Connect(key interface{}, processor FProcessor)
	GetRoute(key interface{}) []FProcessor
	GetRouteKeys() []interface{}
	Trigger(from interface{}, key interface{}, value interface{}) interface{}
}

var IRouteType = reflect.TypeOf((*IRoute)(nil)).Elem()

type IRouteDelegate interface {
	IRoute
	SetDelegate(raw IRoute)
}

var IRouteDelegateType = reflect.TypeOf((*IRouteDelegate)(nil)).Elem()
