package controller

type IControlPointGroup interface {
	GetControlPoint(key interface{}) *ControlPoint
	GetAllControlPoints() map[interface{}]*ControlPoint
	AddControlPoint(key interface{}, point *ControlPoint)
}

type controlPointGroup struct {
	cpoints map[interface{}]*ControlPoint
}

func (this *controlPointGroup) AddControlPoint(key interface{}, point *ControlPoint) {
	this.cpoints[key] = point
}

func (this *controlPointGroup) GetAllControlPoints() map[interface{}]*ControlPoint {
	return this.cpoints
}

func (this *controlPointGroup) GetControlPoint(key interface{}) *ControlPoint {
	return this.cpoints[key]
}

func NewControlPointGroup() IControlPointGroup {
	return &controlPointGroup{
		cpoints: make(map[interface{}]*ControlPoint),
	}
}
