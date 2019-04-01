package gdb

import (
	"github.com/gosrv/goioc"
	"reflect"
)

type IMesageQueueName interface {
	Name() string
}

type IMessageQueueReader interface {
	Pop(num int) []string
}

type IMessageQueueWriter interface {
	Push(msg string) bool
}

type IMessageQueue interface {
	IMesageQueueName
	IMessageQueueReader
	IMessageQueueWriter
}

var IMessageQueueType = reflect.TypeOf((*IMessageQueue)(nil)).Elem()

type IMessageQueueFactory interface {
	gioc.IPriority
	GetMessageQueue(group, id string) IMessageQueue
}
