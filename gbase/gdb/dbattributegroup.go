package gdb

import (
	"github.com/gosrv/goioc"
	"reflect"
)

type IDBAttributeGroup interface {
	CasSetAttribute(key string, oldValue string, newValue string) bool
	GetAttribute(key string) (string, error)
	SetAttribute(key string, value string) error
	SetAttributes(values map[string]interface{}) error
}

var IDBAttributeGroupType = reflect.TypeOf((*IDBAttributeGroup)(nil)).Elem()

type IDBAttributeGroupFactory interface {
	gioc.IPriority
	GetAttributeGroup(group, id string) IDBAttributeGroup
}
