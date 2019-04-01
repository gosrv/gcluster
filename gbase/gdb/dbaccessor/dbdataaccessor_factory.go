package dbaccessor

import "github.com/gosrv/gcluster/gbase/gdb"

type DBDataAccessorFactory struct {
	messageQueueFactory     gdb.IMessageQueueFactory
	attributeGroupFactories []gdb.IDBAttributeGroupFactory
}

func NewDBDataAccessorFactory(messageQueueFactory gdb.IMessageQueueFactory, attributeGroupFactories []gdb.IDBAttributeGroupFactory) *DBDataAccessorFactory {
	return &DBDataAccessorFactory{messageQueueFactory: messageQueueFactory, attributeGroupFactories: attributeGroupFactories}
}

func (this *DBDataAccessorFactory) GetMessageQueue(group, id string) gdb.IMessageQueue {
	return this.messageQueueFactory.GetMessageQueue(group, id)
}

func (this *DBDataAccessorFactory) GetDataAccessor(group, id string) *DBDataAccessor {
	groups := make([]gdb.IDBAttributeGroup, 0, len(this.attributeGroupFactories))
	for _, factory := range this.attributeGroupFactories {
		groups = append(groups, factory.GetAttributeGroup(group, id))
	}
	return NewDBDataAccessor(groups)
}
