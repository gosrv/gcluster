package controller

import (
	"github.com/gosrv/goioc"
	"reflect"
)

type IControlPointGroupMgr interface {
	GetControlPointGroup(group string) IControlPointGroup
	GetControlPoint(group string, tp reflect.Type) *ControlPoint
	GetAllControlGroup() map[string]IControlPointGroup
}

type BeanControlPointGroupMgr struct {
	Controllers []IController `bean:""`
	groups      map[string]IControlPointGroup
}

func (this *BeanControlPointGroupMgr) BeanAfterTagProcess(tagProcessor gioc.ITagProcessor, beanContainer gioc.IBeanContainer) {
	if tagProcessor.TagProcessorName() != gioc.BeanTagProcessor {
		return
	}
	for _, ctrl := range this.Controllers {
		groupName := ctrl.Group()
		cgroup := this.groups[groupName]
		if cgroup == nil {
			cgroup = NewControlPointGroup()
			this.groups[groupName] = cgroup
		}
		ctrl.Finder().ControlPointFind(ctrl, cgroup)
	}
}

func NewBeanControlPointGroupMgr() IControlPointGroupMgr {
	return &BeanControlPointGroupMgr{
		groups: make(map[string]IControlPointGroup),
	}
}

func (this *BeanControlPointGroupMgr) GetAllControlGroup() map[string]IControlPointGroup {
	return this.groups
}

func (this *BeanControlPointGroupMgr) GetControlPointGroup(group string) IControlPointGroup {
	return this.groups[group]
}

func (this *BeanControlPointGroupMgr) GetControlPoint(group string, tp reflect.Type) *ControlPoint {
	cgroup := this.groups[group]
	if cgroup == nil {
		return nil
	}
	return cgroup.GetControlPoint(tp)
}
