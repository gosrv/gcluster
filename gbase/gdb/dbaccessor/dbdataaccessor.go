package dbaccessor

import (
	"github.com/gosrv/gcluster/gbase/gdb"
	"reflect"
)

type DBDataAccessor struct {
	attributeGroups []gdb.IDBAttributeGroup
}

func NewDBDataAccessor(attributeGroups []gdb.IDBAttributeGroup) *DBDataAccessor {
	return &DBDataAccessor{attributeGroups: attributeGroups}
}

func (this *DBDataAccessor) GetAttributeGroup() gdb.IDBAttributeGroup {
	return this.attributeGroups[0]
}

func (this *DBDataAccessor) GetAttributeGroupByType(pt reflect.Type) gdb.IDBAttributeGroup {
	for _, group := range this.attributeGroups {
		if reflect.TypeOf(group).AssignableTo(pt) {
			return group
		}
	}
	return nil
}

func (this *DBDataAccessor) GetDataLoader(attr string) *gdb.TheDataLoaderChain {
	loader := &gdb.TheDataLoaderChain{}
	for _, group := range this.attributeGroups {
		loaderGroup := group
		loader.AddLoader((gdb.FuncTheDataLoader)(func() (string, error) {
			return loaderGroup.GetAttribute(attr)
		}))
	}
	return loader
}

func (this *DBDataAccessor) GetDataSaver(attr string) *gdb.TheDataSaverChain {
	saver := &gdb.TheDataSaverChain{}
	for _, group := range this.attributeGroups {
		loaderGroup := group
		saver.AddSaver(gdb.NewFuncTheDataSaverWraper(func(value string) error {
			return loaderGroup.SetAttribute(attr, value)
		}))
	}
	return saver
}
