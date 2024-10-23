package gometrics

type Fields struct {
	// CPU
	CpuCount       int64 `json:"cpu.count"`
	GoroutineCount int64 `json:"cpu.goroutines"`
	CgoCalls       int64 `json:"cpu.cgo_calls"`

	CpuUsageTotal  uint64 `json:"cpu.useTotal,omitempty"`
	CpuUsageUser   uint64 `json:"cpu.useUser,omitempty"`
	CpuUsageSystem uint64 `json:"cpu.useSystem,omitempty"`
	CpuUsageIdle   uint64 `json:"cpu.useIdle,omitempty"`
	CpuUsageNice   uint64 `json:"cpu.useNice,omitempty"`
	CpuUsageIoWait uint64 `json:"cpu.useIoWait,omitempty"`

	CpuLoadOne     float64 `json:"cpu.loadOne,omitempty"`
	CpuLoadFive    float64 `json:"cpu.loadFive,omitempty"`
	CpuLoadFifteen float64 `json:"cpu.loadFifteen,omitempty"`

	// General
	MemSysTotal uint64 `json:"mem.system.total"`
	MemSysFree  uint64 `json:"mem.system.free"`
	MemSysUsed  uint64 `json:"mem.system.used"`

	Alloc      uint64 `json:"mem.alloc"`
	TotalAlloc uint64 `json:"mem.total"`
	Sys        uint64 `json:"mem.sys"`
	Lookups    uint64 `json:"mem.lookups"`
	Mallocs    uint64 `json:"mem.malloc"`
	Frees      uint64 `json:"mem.frees"`

	// Heap
	HeapAlloc    uint64 `json:"mem.heap.alloc"`
	HeapSys      uint64 `json:"mem.heap.sys"`
	HeapIdle     uint64 `json:"mem.heap.idle"`
	HeapInuse    uint64 `json:"mem.heap.inuse"`
	HeapReleased uint64 `json:"mem.heap.released"`
	HeapObjects  uint64 `json:"mem.heap.objects"`

	// Stack
	StackInuse  uint64 `json:"mem.stack.inuse"`
	StackSys    uint64 `json:"mem.stack.sys"`
	MSpanInuse  uint64 `json:"mem.stack.mspan_inuse"`
	MSpanSys    uint64 `json:"mem.stack.mspan_sys"`
	MCacheInuse uint64 `json:"mem.stack.mcache_inuse"`
	MCacheSys   uint64 `json:"mem.stack.mcache_sys"`

	OtherSys uint64 `json:"mem.othersys"`

	// GC
	GCSys         uint64  `json:"mem.gc.sys"`
	NextGC        uint64  `json:"mem.gc.next"`
	LastGC        uint64  `json:"mem.gc.last"`
	PauseTotalNs  uint64  `json:"mem.gc.pause_total"`
	PauseNs       uint64  `json:"mem.gc.pause"`
	NumGC         uint64  `json:"mem.gc.count"`
	GCCPUFraction float64 `json:"mem.gc.cpu_fraction"`

	Goarch  string `json:"-"`
	Goos    string `json:"-"`
	Version string `json:"-"`
}

func (f Fields) ToMap() map[string]any {
	return map[string]any{
		"CpuCount":       f.CpuCount,
		"GoroutineCount": f.GoroutineCount,
		"CgoCalls":       f.CgoCalls,

		"CpuUsageTotal":  f.CpuUsageTotal,
		"CpuUsageUser":   f.CpuUsageUser,
		"CpuUsageSystem": f.CpuUsageSystem,
		"CpuUsageIdle":   f.CpuUsageIdle,
		"CpuUsageNice":   f.CpuUsageNice,
		"CpuUsageIoWait": f.CpuUsageIoWait,

		"CpuLoadOne":     f.CpuLoadOne,
		"CpuLoadFive":    f.CpuLoadFive,
		"CpuLoadFifteen": f.CpuLoadFifteen,

		"MemSysTotal": f.MemSysTotal,
		"MemSysFree":  f.MemSysFree,
		"MemSysUsed":  f.MemSysUsed,

		"Alloc":      f.Alloc,
		"TotalAlloc": f.TotalAlloc,
		"Sys":        f.Sys,
		"Lookups":    f.Lookups,
		"Mallocs":    f.Mallocs,
		"Frees":      f.Frees,

		"HeapAlloc":    f.HeapAlloc,
		"HeapSys":      f.HeapSys,
		"HeapIdle":     f.HeapIdle,
		"HeapInuse":    f.HeapInuse,
		"HeapReleased": f.HeapReleased,
		"HeapObjects":  f.HeapObjects,

		"StackInuse":  f.StackInuse,
		"StackSys":    f.StackSys,
		"MSpanInuse":  f.MSpanInuse,
		"MSpanSys":    f.MSpanSys,
		"MCacheInuse": f.MCacheInuse,
		"MCacheSys":   f.MCacheSys,

		"OtherSys": f.OtherSys,

		// GC
		"GCSys":         f.GCSys,
		"NextGC":        f.NextGC,
		"LastGC":        f.LastGC,
		"PauseTotalNs":  f.PauseTotalNs,
		"PauseNs":       f.PauseNs,
		"NumGC":         f.NumGC,
		"GCCPUFraction": f.GCCPUFraction,
	}
}
