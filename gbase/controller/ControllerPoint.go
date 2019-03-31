package controller

import "reflect"

type ControlPoint struct {
	Bean        interface{}
	Method      reflect.Value
	ParamTypes  []reflect.Type
	ReturnTypes []reflect.Type
	Controller  IController
}
