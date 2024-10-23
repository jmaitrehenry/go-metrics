package gometrics

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	MaxUint64 = ^uint64(0)
)

func GetCPUUsage() (CPUUsage, error) {
	usage := CPUUsage{}
	contents, err := os.ReadFile("/proc/stat")
	if err != nil {
		return usage, err
	}

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if len(line) > 4 && string(line[0:4]) == "cpu " {
			fields := strings.Fields(string(line))
			usage.User, _ = strconv.ParseUint(fields[1], 10, 64)
			usage.Nice, _ = strconv.ParseUint(fields[2], 10, 64)
			usage.Sys, _ = strconv.ParseUint(fields[3], 10, 64)
			usage.Idle, _ = strconv.ParseUint(fields[4], 10, 64)
			usage.Wait, _ = strconv.ParseUint(fields[5], 10, 64)
			return usage, nil
		}
	}

	return usage, nil
}

func GetLoadAverage() (LoadAverage, error) {
	loadAverage := LoadAverage{}
	line, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return loadAverage, err
	}

	fields := strings.Fields(string(line))
	loadAverage.One, _ = strconv.ParseFloat(fields[0], 64)
	loadAverage.Five, _ = strconv.ParseFloat(fields[1], 64)
	loadAverage.Fifteen, _ = strconv.ParseFloat(fields[2], 64)

	return loadAverage, nil
}

func GetMemoryUsage() (MemUsage, error) {
	memUsage := MemUsage{}
	var available uint64 = MaxUint64
	var buffers, cached uint64
	table := map[string]*uint64{
		"MemTotal":     &memUsage.Total,
		"MemFree":      &memUsage.Free,
		"MemAvailable": &available,
		"Buffers":      &buffers,
		"Cached":       &cached,
	}

	contents, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return memUsage, err
	}

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fields := strings.Split(string(line), ":")

		if ptr := table[fields[0]]; ptr != nil {
			num := strings.TrimLeft(fields[1], " ")
			val, err := strconv.ParseUint(strings.Fields(num)[0], 10, 64)
			if err == nil {
				*ptr = val * 1024
			}
		}
	}

	if available == MaxUint64 {
		memUsage.Free = memUsage.Free + buffers + cached
	} else {
		memUsage.Free = available
	}

	memUsage.Used = memUsage.Total - memUsage.Free

	return memUsage, nil
}
