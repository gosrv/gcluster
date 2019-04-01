package datasync

/**
 * 每一个复合结构继承MarkTags用于标记标志用
 */

type ItemIter func(k interface{}, v interface{})
type ItemStatusIter func(k interface{}, v interface{}, dirty bool)

type IDirtyContainer interface {
	Get(key interface{}) interface{}
	Set(key interface{}, value interface{})
	Foreach(iter ItemIter)
	ForeachStatus(iter ItemStatusIter)
	ForeachDirty(iter ItemIter)
	Size() int
	Clear()
}
