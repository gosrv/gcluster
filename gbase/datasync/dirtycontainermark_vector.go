package datasync

import (
	"github.com/gosrv/goioc/util"
	"math"
)

const (
	int64Byte = 8
)

type DirtyContainerMarkVector struct {
	allDirty        bool
	childContainers []IDirtyContainerMark
	status          []uint64
	// 父节点
	parent IDirtyContainerMark
	// 在父节点中的索引
	idxInParent interface{}
}

func NewDirtyContainerMarkVector() *DirtyContainerMarkVector {
	return &DirtyContainerMarkVector{
		allDirty: true,
		status:   make([]uint64, 1, 1),
	}
}

func (this *DirtyContainerMarkVector) Init(parent IDirtyContainerMark, idxInParent interface{}) {
	util.Assert(this.parent == nil && this.idxInParent == nil, "")
	this.parent = parent
	this.idxInParent = idxInParent
	this.parent.SetChildContainer(this.idxInParent, this)
}

func (this *DirtyContainerMarkVector) Uninit() {
	this.parent.SetChildContainer(this.idxInParent, nil)
	this.parent = nil
	this.idxInParent = nil
}

func (this *DirtyContainerMarkVector) SetChildContainer(idx interface{}, childContainer IDirtyContainerMark) {
	realIdx := idx.(int)
	if childContainer == nil {
		if realIdx < len(this.childContainers) {
			this.childContainers[realIdx] = nil
		}
	} else {
		util.Assert(realIdx >= 0 && realIdx < math.MaxInt16, "")
		if realIdx >= len(this.childContainers) {
			this.childContainers = append(this.childContainers,
				make([]IDirtyContainerMark, realIdx-len(this.childContainers)+1)...)
		}
		old := this.childContainers[realIdx]
		this.childContainers[realIdx] = childContainer
		util.Assert(old == nil, "")
	}
}

func (this *DirtyContainerMarkVector) MarkDirtyUp(key interface{}) {
	if key != nil {
		realIdx := key.(int)
		if !this.SetDirty(realIdx) {
			return
		}
	}
	if this.parent != nil {
		this.parent.MarkDirtyUp(this.idxInParent)
	}
}

func (this *DirtyContainerMarkVector) IsAllDirty() bool {
	return this.allDirty
}

func (this *DirtyContainerMarkVector) MarkAllDirty() {
	this.allDirty = true
	this.MarkDirtyUp(nil)
}

func (this *DirtyContainerMarkVector) expand(tosize int) {
	util.Assert(tosize*8 < math.MaxInt16, "")
	if len(this.status) >= tosize {
		return
	}
	for i := len(this.status); i <= tosize; i++ {
		this.status = append(this.status, 0)
	}
}

func (this *DirtyContainerMarkVector) SetDirty(idx interface{}) bool {
	realIdx := idx.(int)
	idxStatus := realIdx / int64Byte
	offsetStatus := realIdx % int64Byte
	if idxStatus >= len(this.status) {
		newStatusSize := int(math.Max(float64(len(this.status)*2), float64(idxStatus+1.0)))
		this.expand(newStatusSize)
	}

	if ((this.status[idxStatus] >> uint(offsetStatus)) & 0x01) != 0 {
		return false
	}
	this.status[idxStatus] |= uint64(1) << uint(offsetStatus)
	return true
}

func (this *DirtyContainerMarkVector) IsDirty(idx interface{}) bool {
	if this.allDirty {
		return true
	}
	realIdx := idx.(int)
	if realIdx < 0 || realIdx >= len(this.status)*int64Byte {
		return false
	}
	idxStatus := realIdx / int64Byte
	offsetStatus := realIdx % int64Byte
	return ((this.status[idxStatus] >> uint64(offsetStatus)) & uint64(1)) != 0
}

func (this *DirtyContainerMarkVector) ClearDirty() {
	for i, v := range this.childContainers {
		if !this.IsDirty(i) {
			continue
		}
		if v != nil {
			v.ClearDirty()
		}
	}
	for i := 0; i < len(this.status); i++ {
		this.status[i] = 0
	}
	this.allDirty = false
}
