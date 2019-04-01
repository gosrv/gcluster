package gnet

import (
	"reflect"
)

const (
	ScopeSession = 0
	ScopeRequest = 1
	ScopeAll     = 2
)

type ISessionCtx interface {
	SetAttribute(scope int, key interface{}, val interface{})
	GetAttribute(scope int, key interface{}) interface{}
	Set(key interface{}, val interface{})
	Get(key interface{}) interface{}
	GetByType(tp reflect.Type) interface{}
	GetAllByType(tp reflect.Type) []interface{}
	Clear(scope int)
}

var ISessionCtxType = reflect.TypeOf((*ISessionCtx)(nil)).Elem()

type sessionCtx struct {
	sessionAttributes [ScopeAll]map[interface{}]interface{}
}

func NewSessionCtx() ISessionCtx {
	ctx := &sessionCtx{}
	for idx := range ctx.sessionAttributes {
		ctx.sessionAttributes[idx] = make(map[interface{}]interface{})
	}
	return ctx
}

func (this *sessionCtx) SetAttribute(scope int, key interface{}, val interface{}) {
	this.sessionAttributes[scope][key] = val
}

func (this *sessionCtx) GetAttribute(scope int, key interface{}) interface{} {
	return this.sessionAttributes[scope][key]
}

func (this *sessionCtx) Set(key interface{}, val interface{}) {
	this.SetAttribute(ScopeSession, key, val)
}

func (this *sessionCtx) Get(key interface{}) interface{} {
	for _, attrs := range this.sessionAttributes {
		val := attrs[key]
		if val != nil {
			return val
		}
	}
	return nil
}

func (this *sessionCtx) Clear(scope int) {
	if scope == ScopeAll {
		for idx := range this.sessionAttributes {
			this.sessionAttributes[idx] = make(map[interface{}]interface{})
		}
	} else {
		this.sessionAttributes[scope] = make(map[interface{}]interface{})
	}
}

func (this *sessionCtx) GetByType(tp reflect.Type) interface{} {
	// 先进行精确查找
	ins := this.Get(tp)
	if ins != nil && reflect.TypeOf(ins).AssignableTo(tp) {
		return ins
	}

	// 遍历查找
	for _, attrs := range this.sessionAttributes {
		for _, v := range attrs {
			if reflect.TypeOf(v).AssignableTo(tp) {
				return v
			}
		}
	}
	return nil
}

func (this *sessionCtx) GetAllByType(tp reflect.Type) []interface{} {
	var all []interface{}
	// 遍历查找
	for _, attrs := range this.sessionAttributes {
		for _, v := range attrs {
			if reflect.TypeOf(v).AssignableTo(tp) {
				all = append(all, v)
			}
		}
	}
	return all
}
