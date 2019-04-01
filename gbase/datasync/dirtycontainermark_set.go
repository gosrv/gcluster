package datasync

import "github.com/gosrv/goioc/util"

type DirtyContainerMarkSet struct {
	// 脏key
	dirtys map[interface{}]bool
	// 所有的key都脏
	allDirty bool
	// 父节点
	parent IDirtyContainerMark
	// 在父节点中的索引
	idxInParent     interface{}
	childContainers map[interface{}]IDirtyContainerMark
}

func NewDirtyContainerMarkSet() IDirtyContainerMark {
	return &DirtyContainerMarkSet{
		allDirty: true,
	}
}

func (this *DirtyContainerMarkSet) Init(parent IDirtyContainerMark, idxInParent interface{}) {
	util.Assert(this.parent == nil && this.idxInParent == nil, "")
	this.parent = parent
	this.idxInParent = idxInParent
	this.parent.SetChildContainer(this.idxInParent, this)
}

func (this *DirtyContainerMarkSet) Uninit() {
	this.parent.SetChildContainer(this.idxInParent, nil)
	this.parent = nil
	this.idxInParent = nil
}

func (this *DirtyContainerMarkSet) SetChildContainer(idx interface{}, childContainer IDirtyContainerMark) {
	if childContainer == nil {
		delete(this.childContainers, idx)
	} else {
		_, hasold := this.childContainers[idx]
		util.Assert(!hasold, "")
		this.childContainers[idx] = childContainer
	}
}

func (this *DirtyContainerMarkSet) MarkDirtyUp(key interface{}) {
	if key != nil {
		if !this.SetDirty(key) {
			return
		}
	}
	if this.parent != nil {
		this.parent.MarkDirtyUp(this.idxInParent)
	}
}

func (this *DirtyContainerMarkSet) MarkAllDirty() {
	this.allDirty = true
	this.MarkDirtyUp(nil)
}

func (this *DirtyContainerMarkSet) IsAllDirty() bool {
	return this.allDirty
}

func (this *DirtyContainerMarkSet) SetDirty(key interface{}) bool {
	_, oldok := this.dirtys[key]
	this.dirtys[key] = true
	return oldok
}

func (this *DirtyContainerMarkSet) IsDirty(key interface{}) bool {
	_, ok := this.dirtys[key]
	return this.allDirty || ok
}

func (this *DirtyContainerMarkSet) ClearDirty() {
	for k, v := range this.childContainers {
		if this.IsDirty(k) {
			v.ClearDirty()
		}
	}
	if len(this.dirtys) > 0 {
		this.dirtys = make(map[interface{}]bool)
	}
	this.allDirty = false
}
