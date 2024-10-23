package gometrics

/*
#include <stdlib.h>
#include <sys/sysctl.h>
#include <sys/mount.h>
#include <mach/mach_init.h>
#include <mach/mach_host.h>
#include <mach/host_info.h>
#include <libproc.h>
#include <mach/processor_info.h>
#include <mach/vm_map.h>
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"
)

func GetCPUUsage() (CPUUsage, error) {
	usage := CPUUsage{}
	var count C.mach_msg_type_number_t = C.HOST_CPU_LOAD_INFO_COUNT
	var cpuload C.host_cpu_load_info_data_t

	status := C.host_statistics(C.host_t(C.mach_host_self()),
		C.HOST_CPU_LOAD_INFO,
		C.host_info_t(unsafe.Pointer(&cpuload)),
		&count)

	if status != C.KERN_SUCCESS {
		return usage, fmt.Errorf("host_statistics error=%d", status)
	}

	usage.User = uint64(cpuload.cpu_ticks[C.CPU_STATE_USER])
	usage.Sys = uint64(cpuload.cpu_ticks[C.CPU_STATE_SYSTEM])
	usage.Idle = uint64(cpuload.cpu_ticks[C.CPU_STATE_IDLE])
	usage.Nice = uint64(cpuload.cpu_ticks[C.CPU_STATE_NICE])

	return usage, nil
}

func GetLoadAverage() (LoadAverage, error) {
	avg := []C.double{0, 0, 0}

	C.getloadavg(&avg[0], C.int(len(avg)))

	return LoadAverage{
		One:     float64(avg[0]),
		Five:    float64(avg[1]),
		Fifteen: float64(avg[2]),
	}, nil
}

func GetMemoryUsage() (MemUsage, error) {
	var vmstat C.vm_statistics_data_t
	memUsage := MemUsage{}

	if err := sysctlbyname("hw.memsize", &memUsage.Total); err != nil {
		return memUsage, err
	}

	if err := vm_info(&vmstat); err != nil {
		return memUsage, err
	}

	kern := uint64(vmstat.inactive_count) << 12
	memUsage.Free = (uint64(vmstat.free_count) << 12) + kern
	memUsage.Used = memUsage.Total - memUsage.Free - kern

	return memUsage, nil
}

func sysctlbyname(name string, data interface{}) (err error) {
	val, err := syscall.Sysctl(name)
	if err != nil {
		return err
	}

	buf := []byte(val)

	switch v := data.(type) {
	case *uint64:
		*v = *(*uint64)(unsafe.Pointer(&buf[0]))
		return
	}

	bbuf := bytes.NewBuffer([]byte(val))
	return binary.Read(bbuf, binary.LittleEndian, data)
}

func vm_info(vmstat *C.vm_statistics_data_t) error {
	var count C.mach_msg_type_number_t = C.HOST_VM_INFO_COUNT

	status := C.host_statistics(
		C.host_t(C.mach_host_self()),
		C.HOST_VM_INFO,
		C.host_info_t(unsafe.Pointer(vmstat)),
		&count)

	if status != C.KERN_SUCCESS {
		return fmt.Errorf("host_statistics=%d", status)
	}

	return nil
}
