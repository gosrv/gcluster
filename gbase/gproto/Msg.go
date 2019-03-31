package gproto

import "reflect"

type IMsgProcessor interface {
	ProcessMsg(msg interface{}) bool
}

var IMsgProcessorType = reflect.TypeOf((*IMsgProcessor)(nil)).Elem()
