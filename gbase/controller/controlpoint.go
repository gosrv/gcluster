package controller

import "reflect"

/**
控制点，每个路由函数是一个控制点
*/
type ControlPoint struct {
	Bean        interface{}
	Method      reflect.Value
	ParamTypes  []reflect.Type
	ReturnTypes []reflect.Type
	Controller  IController
}
