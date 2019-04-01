package controller

import (
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gbase/gnet"
	"reflect"
)

type IControlPointFinder interface {
	ControlPointFind(bean IController, group IControlPointGroup)
}

type FuncTypeControlPointFinder func(bean IController, group IControlPointGroup)

func (this FuncTypeControlPointFinder) ControlPointFind(bean IController, group IControlPointGroup) {
	this(bean, group)
}

func typeControlPointFinder(bean IController, group IControlPointGroup) {
	rval := reflect.ValueOf(bean)
	if rval.Kind() != reflect.Ptr || rval.Elem().Kind() == reflect.Ptr {
		panic("bean only support pointer to struct")
	}
	for i := 0; i < rval.NumMethod(); i++ {
		method := rval.Method(i)
		if method.Type().NumIn() < 2 {
			continue
		}
		// 入参至少两个，第一个是ISessionCtx，第二个是消息
		if method.Type().In(0) != reflect.TypeOf((*gnet.ISessionCtx)(nil)).Elem() {
			continue
		}
		// 返回参数最多1个
		if method.Type().NumOut() >= 2 {
			continue
		}
		keyt := method.Type().In(1)
		if group.GetControlPoint(keyt) != nil {
			glog.Panic("duplicate control point %v in %v and %v", keyt, reflect.TypeOf(bean),
				reflect.TypeOf(group.GetControlPoint(keyt).Bean))
		}

		ctlPoint := &ControlPoint{
			Bean:        bean,
			Method:      method,
			ParamTypes:  make([]reflect.Type, method.Type().NumIn(), method.Type().NumIn()),
			ReturnTypes: make([]reflect.Type, method.Type().NumIn(), method.Type().NumIn()),
			Controller:  bean,
		}
		for j := 0; j < method.Type().NumIn(); j++ {
			ctlPoint.ParamTypes[j] = method.Type().In(j)
		}
		for j := 0; j < method.Type().NumOut(); j++ {
			ctlPoint.ReturnTypes[j] = method.Type().Out(j)
		}
		group.AddControlPoint(keyt, ctlPoint)
	}
}
