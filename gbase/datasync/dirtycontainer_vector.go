package datasync

import (
	"github.com/gosrv/goioc/util"
	"math"
)

type DirtyContainerVector struct {
	DirtyContainerMarkVector
	listDatas []interface{}
}

func NewDirtyContainerVector() *DirtyContainerVector {
	return &DirtyContainerVector{
		DirtyContainerMarkVector: DirtyContainerMarkVector{
			allDirty: true,
		},
	}
}

func (this *DirtyContainerVector) Get(idx interface{}) interface{} {
	realIdx := idx.(int)
	if realIdx < 0 || realIdx >= len(this.listDatas) {
		return nil
	}

	return this.listDatas[realIdx]
}

func (this *DirtyContainerVector) Set(idx interface{}, val interface{}) {
	realIdx := idx.(int)
	util.Assert(realIdx < math.MaxInt16, "")

	if realIdx >= len(this.listDatas) {
		this.listDatas = this.listDatas[:realIdx+1]
	}

	old := this.listDatas[realIdx]
	this.listDatas[realIdx] = val
	if old != nil {
		if dirtyContainerMark, ok := old.(IDirtyContainerMark); ok {
			dirtyContainerMark.Uninit()
		}
	}

	if val != nil {
		if dirtyContainerMark, ok := val.(IDirtyContainerMark); ok {
			container := dirtyContainerMark
			container.Init(this, idx)
			container.MarkAllDirty()
		}
	}

	this.MarkDirtyUp(idx)
}

func (this *DirtyContainerVector) Foreach(iter ItemIter) {
	for k, v := range this.listDatas {
		iter(k, v)
	}
}

func (this *DirtyContainerVector) ForeachStatus(iter ItemStatusIter) {
	for k, v := range this.listDatas {
		iter(k, v, this.IsDirty(k))
	}
}

func (this *DirtyContainerVector) ForeachDirty(iter ItemIter) {
	for k, v := range this.listDatas {
		if !this.IsDirty(k) {
			continue
		}
		iter(k, v)
	}
}

func (this *DirtyContainerVector) Size() int {
	return len(this.listDatas)
}

func (this *DirtyContainerVector) Clear() {
	for i, k := range this.listDatas {
		if k != nil {
			if dirtyContainerMark, ok := k.(IDirtyContainerMark); ok {
				dirtyContainerMark.Uninit()
			}
		}
		this.listDatas[i] = nil
	}

	if len(this.listDatas) > 0 {
		this.MarkAllDirty()
		this.MarkDirtyUp(nil)
	}
}

func (this *DirtyContainerVector) clearDirty() {
	cutIdx := len(this.listDatas)
	for cutIdx := len(this.listDatas) - 1; cutIdx >= 0; cutIdx-- {
		if this.listDatas[cutIdx] != nil {
			break
		}
	}
	cutIdx++
	if cutIdx < len(this.listDatas) {
		this.listDatas = this.listDatas[:cutIdx]
	}

	this.DirtyContainerMarkVector.ClearDirty()
}
