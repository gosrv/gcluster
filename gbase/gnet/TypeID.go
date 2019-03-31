package gnet

import (
	"errors"
	"fmt"
	"github.com/gosrv/goioc/util"
	"reflect"
)

type ITypeID interface {
	ID2Type(id int) (reflect.Type, error)
	Type2ID(tp reflect.Type) (int, error)
	AddIDType(id int, tp reflect.Type) error
}

type typeID struct {
	id2typeMap map[int]reflect.Type
	type2idMap map[reflect.Type]int
}

func NewTypeID() ITypeID {
	return &typeID{
		id2typeMap: make(map[int]reflect.Type),
		type2idMap: make(map[reflect.Type]int),
	}
}

func (this *typeID) AddIDType(id int, tp reflect.Type) error {
	util.Assert(tp != nil, fmt.Sprintf("id [%v] has nil type", id))
	_, hasid := this.id2typeMap[id]
	if hasid {
		return errors.New("duplicate id type map")
	}
	_, hastp := this.type2idMap[tp]
	if hastp {
		return errors.New("duplicate id type map")
	}

	this.id2typeMap[id] = tp
	this.type2idMap[tp] = id
	return nil
}

func (this *typeID) ID2Type(id int) (reflect.Type, error) {
	val, ok := this.id2typeMap[id]
	if !ok {
		return nil, errors.New("no item")
	}
	return val, nil
}

func (this *typeID) Type2ID(tp reflect.Type) (int, error) {
	val, ok := this.type2idMap[tp]
	if !ok {
		return -1, errors.New("no item")
	}
	return val, nil
}
