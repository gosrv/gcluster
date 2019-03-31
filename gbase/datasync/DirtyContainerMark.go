package datasync

import "reflect"

type IDirtyContainerMark interface {
	// 设置父容器关联
	Init(parent IDirtyContainerMark, idxInParent interface{})
	SetChildContainer(idx interface{}, childContainer IDirtyContainerMark)
	// 取消与父容器关联
	Uninit()
	// 向上（父）标记脏
	MarkDirtyUp(key interface{})
	// 标记所有内容脏
	MarkAllDirty()
	// 是否所有内容脏
	IsAllDirty() bool
	// 设置某个key脏
	SetDirty(key interface{}) bool
	// 某个key是否脏
	IsDirty(key interface{}) bool
	// 清除脏标记
	ClearDirty()
}

var IDirtyContainerMarkType = reflect.TypeOf((*IDirtyContainerMark)(nil)).Elem()
