package datasync

import (
	"github.com/gosrv/goioc/util"
)

type DirtyValue struct {
	value interface{}
	dirty bool
}

func NewDirtyValue(value interface{}, dirty bool) *DirtyValue {
	return &DirtyValue{value: value, dirty: dirty}
}

type DirtyContainerMap struct {
	allDirty bool
	// 父节点
	parent IDirtyContainerMark
	// 在父节点中的索引
	idxInParent interface{}
	datas       map[interface{}]*DirtyValue
	dirtys      map[interface{}]*DirtyValue
}

func NewDirtyContainerMap() *DirtyContainerMap {
	return &DirtyContainerMap{
		allDirty: true,
		datas:    make(map[interface{}]*DirtyValue),
		dirtys:   make(map[interface{}]*DirtyValue),
	}
}

func (this *DirtyContainerMap) Get(key interface{}) interface{} {
	dv, ok := this.datas[key]
	if !ok || dv == nil {
		return nil
	}
	return dv.value
}

func (this *DirtyContainerMap) Set(key interface{}, val interface{}) {
	dv := NewDirtyValue(val, true)
	old := this.datas[key]
	this.datas[key] = dv
	if old != nil && old.value != nil {
		if dirtyContainerMark, ok := old.value.(IDirtyContainerMark); ok {
			dirtyContainerMark.Uninit()
		}
	}
	this.dirtys[key] = dv

	if val != nil {
		if dirtyContainerMark, ok := val.(IDirtyContainerMark); ok {
			dirtyContainerMark.Init(this, key)
			dirtyContainerMark.MarkAllDirty()
		}
	}

	this.MarkDirtyUp(key)
}

func (this *DirtyContainerMap) Foreach(iter ItemIter) {
	for k, v := range this.datas {
		iter(k, v.value)
	}
}

func (this *DirtyContainerMap) ForeachStatus(iter ItemStatusIter) {
	if this.allDirty {
		for k, v := range this.datas {
			iter(k, v.value, true)
		}
	} else {
		for k, v := range this.datas {
			iter(k, v.value, v.dirty)
		}
	}

}

func (this *DirtyContainerMap) ForeachDirty(iter ItemIter) {
	if this.allDirty {
		for k, v := range this.datas {
			iter(k, v.value)
		}
	} else {
		for k, v := range this.dirtys {
			iter(k, v.value)
		}
	}
}

func (this *DirtyContainerMap) Size() int {
	return len(this.datas)
}

func (this *DirtyContainerMap) Clear() {
	// 这里只能置空，不能删除，删除要等cleardirty时进行，不然无法知道哪些值被clear了
	for _, val := range this.datas {
		if val != nil && val.value != nil {
			if dirtyContainerMark, ok := val.value.(IDirtyContainerMark); ok {
				dirtyContainerMark.Uninit()
			}
		}
		val.value = nil
	}

	if len(this.datas) > 0 {
		this.MarkAllDirty()
	}
}

func (this *DirtyContainerMap) Init(parent IDirtyContainerMark, idxInParent interface{}) {
	util.Assert(this.parent == nil && this.idxInParent == nil, "")
	this.parent = parent
	this.idxInParent = idxInParent
	this.parent.SetChildContainer(this.idxInParent, this)
}

func (this *DirtyContainerMap) Uninit() {
	this.parent.SetChildContainer(this.idxInParent, nil)
	this.parent = nil
	this.idxInParent = nil
}

func (this *DirtyContainerMap) SetChildContainer(idx interface{}, childContainer IDirtyContainerMark) {
}

func (this *DirtyContainerMap) MarkDirtyUp(key interface{}) {
	if key != nil {
		if !this.SetDirty(key) {
			return
		}
	}
	if this.parent != nil {
		this.parent.MarkDirtyUp(this.idxInParent)
	}
}

func (this *DirtyContainerMap) MarkAllDirty() {
	this.allDirty = true
	this.MarkDirtyUp(nil)
}

func (this *DirtyContainerMap) IsAllDirty() bool {
	return this.allDirty
}

func (this *DirtyContainerMap) SetDirty(key interface{}) bool {
	dv := this.datas[key]
	if dv != nil {
		dv.dirty = true
	}
	_, ok := this.dirtys[key]
	this.dirtys[key] = dv
	return !ok
}

func (this *DirtyContainerMap) IsDirty(key interface{}) bool {
	_, ok := this.dirtys[key]
	return this.allDirty || ok
}

func (this *DirtyContainerMap) ClearDirty() {
	for k, v := range this.dirtys {
		if v == nil && v.value == nil {
			delete(this.dirtys, k)
		} else if !this.allDirty {
			if dirtyContainerMark, ok := v.value.(IDirtyContainerMark); ok {
				dirtyContainerMark.ClearDirty()
			}
		}
	}

	if this.allDirty {
		for k, v := range this.datas {
			if v == nil || v.value == nil {
				delete(this.datas, k)
			} else if dirtyContainerMark, ok := v.value.(IDirtyContainerMark); ok {
				dirtyContainerMark.ClearDirty()
			}
		}
	}

	if len(this.dirtys) > 0 {
		this.dirtys = make(map[interface{}]*DirtyValue)
	}
	this.allDirty = false
}
