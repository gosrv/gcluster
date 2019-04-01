package ghttp

import (
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/goioc/util"
	"reflect"
)

const (
	InjectParam  = "p"
	InjectForm   = "f"
	InjectHeader = "h"
	InjectCookie = "c"
)

func injectParam(val reflect.Value, name string, param *HttpParam) bool {
	pv, ok := param.Params[name]
	if !ok || len(pv) == 0 {
		return false
	}
	if val.Kind() == reflect.Slice {
		val.Set(reflect.ValueOf(pv))
	} else {
		val.Set(reflect.ValueOf(pv[0]))
	}
	return true
}

func injectForm(val reflect.Value, name string, param *HttpForm) bool {
	pv, ok := param.Params[name]
	if !ok || len(pv) == 0 {
		return false
	}
	if val.Kind() == reflect.Slice {
		val.Set(reflect.ValueOf(pv))
	} else {
		val.Set(reflect.ValueOf(pv[0]))
	}
	return true
}

func injectHeader(val reflect.Value, name string, param *HttpHeader) bool {
	pv, ok := param.Headers[name]
	if !ok || len(pv) == 0 {
		return false
	}
	if val.Kind() == reflect.Slice {
		val.Set(reflect.ValueOf(pv))
	} else {
		val.Set(reflect.ValueOf(pv[0]))
	}
	return true
}

func injectCookie(val reflect.Value, name string, param *HttpCookie) bool {
	pv, ok := param.ParamSingle[name]
	if !ok {
		return false
	}
	val.Set(reflect.ValueOf(pv))
	return true
}

func ParamInject(val reflect.Value, ctx gnet.ISessionCtx) {
	headers := ctx.Get(reflect.TypeOf((*HttpHeader)(nil))).(*HttpHeader)
	params := ctx.Get(reflect.TypeOf((*HttpParam)(nil))).(*HttpParam)
	cookies := ctx.Get(reflect.TypeOf((*HttpCookie)(nil))).(*HttpCookie)
	forms := ctx.Get(reflect.TypeOf((*HttpForm)(nil))).(*HttpForm)

	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldValue = util.Hack.ValuePatchWrite(fieldValue)
		if name, ok := fieldType.Tag.Lookup(InjectParam); ok {
			_ = injectParam(fieldValue, name, params)
		} else if name, ok := fieldType.Tag.Lookup(InjectForm); ok {
			_ = injectForm(fieldValue, name, forms)
		} else if name, ok := fieldType.Tag.Lookup(InjectHeader); ok {
			_ = injectHeader(fieldValue, name, headers)
		} else if name, ok := fieldType.Tag.Lookup(InjectCookie); ok {
			_ = injectCookie(fieldValue, name, cookies)
		} else {
			_ = injectParam(fieldValue, fieldType.Name, params) ||
				injectForm(fieldValue, fieldType.Name, forms) ||
				injectHeader(fieldValue, fieldType.Name, headers) ||
				injectCookie(fieldValue, fieldType.Name, cookies)
		}
	}
}
