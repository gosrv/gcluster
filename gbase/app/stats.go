package app

import (
	"runtime"
	"time"
)

type appStats struct {
}

var AppStats = &appStats{}

type RuntimeStats struct {
	// 累积GC暂停时间
	PauseTotal int64
	// 最近一次GC暂停时间
	LastPause int64
	// 协程数量
	NumGoroutine int
	// CPU数
	NumCpu int
	// CGO 调用次数
	NumCgoCall int64
}

func (this *appStats) GetMemStats() *runtime.MemStats {
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	return memStats
}

func (this *appStats) GetRuntimeStats() RuntimeStats {
	memStats := this.GetMemStats()
	var stats RuntimeStats
	// 累积GC暂停时间
	stats.PauseTotal = int64(time.Duration(memStats.PauseTotalNs) / time.Microsecond)
	// 最近一次GC暂停时间
	stats.LastPause = int64(time.Duration(memStats.PauseNs[(memStats.NumGC+255)%256]) / time.Microsecond)
	// 协程数量
	stats.NumGoroutine = runtime.NumGoroutine()
	// CGO 调用次数
	stats.NumCgoCall = runtime.NumCgoCall()
	// CPU数
	stats.NumCpu = runtime.NumCPU()

	return stats
}
