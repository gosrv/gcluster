package gdb

import (
	"reflect"
	"time"
)

type IDataExpireable interface {
	SetExpireDuration(duration time.Duration) error
}

var IDataExpireableType = reflect.TypeOf((*IDataExpireable)(nil)).Elem()
