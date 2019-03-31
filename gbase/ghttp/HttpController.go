package ghttp

import (
	"github.com/gosrv/gcluster/gbase/controller"
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gbase/gnet"
	"reflect"
	"strings"
)

const (
	HttpController = "HttpController"
)

type IHttpController interface {
	controller.IController
	ViewRender() IViewRender
}

var IHttpControllerType = reflect.TypeOf((*IHttpController)(nil)).Elem()

type httpController struct {
	group      string
	ctype      string
	finder     controller.IControlPointFinder
	trigger    controller.IControlPointTrigger
	viewRender IViewRender
}

func (this *httpController) Group() string {
	return this.group
}

func (this *httpController) Type() string {
	return this.ctype
}

func (this *httpController) Finder() controller.IControlPointFinder {
	return this.finder
}

func (this *httpController) Trigger() controller.IControlPointTrigger {
	return this.trigger
}

func (this *httpController) ViewRender() IViewRender {
	return this.viewRender
}

func NewHttpRestController(group string) IHttpController {
	return NewHttpController(group, NewRestViewRender())
}

func NewHttpController(group string, viewRender IViewRender) IHttpController {
	return &httpController{
		group:      group,
		ctype:      HttpController,
		finder:     controller.FuncTypeControlPointFinder(httpControlPointFinder),
		trigger:    controller.FuncControlPointTrigger(httpControlPointTrigger),
		viewRender: viewRender,
	}
}

func httpControlPointFinder(bean controller.IController, group controller.IControlPointGroup) {
	rval := reflect.ValueOf(bean)
	rtype := reflect.TypeOf(bean)
	if rval.Kind() != reflect.Ptr || rval.Elem().Kind() == reflect.Ptr {
		panic("bean only support pointer to struct")
	}
	for i := 0; i < rval.NumMethod(); i++ {
		method := rval.Method(i)
		mtype := rtype.Method(i)
		// 入参一个是ISessionCtx
		if method.Type().NumIn() < 1 || method.Type().In(0) != gnet.ISessionCtxType {
			continue
		}

		// 返回参数最多1个
		if method.Type().NumOut() >= 2 {
			continue
		}
		methodName := mtype.Name
		key := "/"
		for k, m := range methodName {
			if m >= 'A' && m <= 'Z' && k != 0 {
				key += "/" + strings.ToLower(string(m))
			} else {
				key += strings.ToLower(string(m))
			}
		}
		if group.GetControlPoint(key) != nil {
			glog.Panic("duplicate control point %v in %v and %v", key, reflect.TypeOf(bean),
				reflect.TypeOf(group.GetControlPoint(key).Bean))
		}

		ctlPoint := &controller.ControlPoint{
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
		group.AddControlPoint(key, ctlPoint)
	}
}

func httpControlPointTrigger(controlPoint *controller.ControlPoint, ctx gnet.ISessionCtx) interface{} {
	params := make([]reflect.Value, len(controlPoint.ParamTypes), len(controlPoint.ParamTypes))
	for i := 0; i < len(params); i++ {
		val := ctx.GetByType(controlPoint.ParamTypes[i])
		if val == nil {
			nval := reflect.New(controlPoint.ParamTypes[i]).Elem()
			params[i] = nval
			for nval.Kind() == reflect.Ptr {
				nval.Set(reflect.New(nval.Type().Elem()))
				nval = nval.Elem()
			}
			// 注入数据
			ParamInject(nval, ctx)
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
