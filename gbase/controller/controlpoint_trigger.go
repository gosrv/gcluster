package controller

import (
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gbase/gnet"
	"reflect"
)

type IControlPointTrigger interface {
	Trigger(controlPoint *ControlPoint, ctx gnet.ISessionCtx) interface{}
}

type FuncControlPointTrigger func(controlPoint *ControlPoint, ctx gnet.ISessionCtx) interface{}

func (this FuncControlPointTrigger) Trigger(controlPoint *ControlPoint, ctx gnet.ISessionCtx) interface{} {
	return this(controlPoint, ctx)
}

func typeControlPointTrigger(controlPoint *ControlPoint, ctx gnet.ISessionCtx) interface{} {
	params := make([]reflect.Value, len(controlPoint.ParamTypes), len(controlPoint.ParamTypes))
	for i := 0; i < len(params); i++ {
		val := ctx.GetByType(controlPoint.ParamTypes[i])
		if val == nil {
			params[i] = reflect.New(controlPoint.ParamTypes[i]).Elem()
		} else {
			switch val.(type) {
			case reflect.Value:
				params[i] = val.(reflect.Value)
			default:
				params[i] = reflect.ValueOf(val)
			}
		}
	}
	reps := controlPoint.Method.Call(params)
	switch len(reps) {
	case 0:
		return nil
	case 1:
		return reps[0].Interface()
	default:
		glog.Panic("invalid return num, expect 1")
	}
	return nil
}
