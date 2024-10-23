package gometrics

import (
	"fmt"
	"runtime"
)

// Collector implements the grabbing of metrics
type Collector struct {
	// Namespace is a component of the fully-qualified name of the Metric
	Namespace string

	// Subsystem is a component of the fully-qualified name of the Metric
	Subsystem string

	// EnableCPU determines whether CPU statistics will be output. Defaults to true.
	EnableCPU bool

	// EnableMem determines whether memory statistics will be output. Defaults to true.
	EnableMem bool
}

func (c *Collector) collectRuntimeInfo() Fields {
	fields := Fields{}

	fields.Goos = runtime.GOOS
	fields.Goarch = runtime.GOARCH
	fields.Version = runtime.Version()

	return fields
}

func (c *Collector) CollectStats() Fields {
	fields := Fields{}

	if c.EnableCPU {
		fields.GoroutineCount = int64(runtime.NumGoroutine())
		fields.CgoCalls = int64(runtime.NumCgoCall())
		fields.CpuCount = int64(runtime.NumCPU())
		cpuStats, err := GetCPUUsage()
		if err != nil {
			fmt.Printf("Error getting CPU usage: %v\n", err)
		} else {
			fields.CpuUsageTotal = cpuStats.User + cpuStats.Sys + cpuStats.Nice
			fields.CpuUsageUser = cpuStats.User
			fields.CpuUsageSystem = cpuStats.Sys
			fields.CpuUsageIdle = cpuStats.Idle
			fields.CpuUsageNice = cpuStats.Nice
			fields.CpuUsageIoWait = cpuStats.Wait
		}

		load, err := GetLoadAverage()
		if err != nil {
			fmt.Printf("Error getting load average: %v\n", err)
		} else {
			fields.CpuLoadOne = load.One
			fields.CpuLoadFive = load.Five
			fields.CpuLoadFifteen = load.Fifteen
		}
	}

	if c.EnableMem {
		m := &runtime.MemStats{}
		runtime.ReadMemStats(m)

		// System
		memUsage, err := GetMemoryUsage()
		if err != nil {
			fmt.Printf("Error getting memory usage: %v\n", err)
		} else {
			fields.MemSysTotal = memUsage.Total
			fields.MemSysFree = memUsage.Free
			fields.MemSysUsed = memUsage.Used
		}

		// General
		fields.Alloc = m.Alloc
		fields.TotalAlloc = m.TotalAlloc
		fields.Sys = m.Sys
		fields.Lookups = m.Lookups
		fields.Mallocs = m.Mallocs
		fields.Frees = m.Frees

		// Heap
		fields.HeapAlloc = m.HeapAlloc
		fields.HeapSys = m.HeapSys
		fields.HeapIdle = m.HeapIdle
		fields.HeapInuse = m.HeapInuse
		fields.HeapReleased = m.HeapReleased
		fields.HeapObjects = m.HeapObjects

		// Stack
		fields.StackInuse = m.StackInuse
		fields.StackSys = m.StackSys
		fields.MSpanInuse = m.MSpanInuse
		fields.MSpanSys = m.MSpanSys
		fields.MCacheInuse = m.MCacheInuse
		fields.MCacheSys = m.MCacheSys

		fields.OtherSys = m.OtherSys

		// Garbage Collector
		fields.GCSys = m.GCSys
		fields.NextGC = m.NextGC
		fields.LastGC = m.LastGC
		fields.PauseTotalNs = m.PauseTotalNs
		fields.PauseNs = m.PauseNs[(m.NumGC+255)%256]
		fields.NumGC = uint64(m.NumGC)
		fields.GCCPUFraction = m.GCCPUFraction
	}

	fields.Goos = runtime.GOOS
	fields.Goarch = runtime.GOARCH
	fields.Version = runtime.Version()

	return fields
}
