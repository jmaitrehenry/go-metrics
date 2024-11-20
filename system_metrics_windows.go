package gometrics

import "errors"

func GetCPUUsage() (CPUUsage, error) {
	return CPUUsage{}, errors.New("Not implemented on Windows")
}

func GetLoadAverage() (LoadAverage, error) {
	return LoadAverage{}, errors.New("Not implemented on Windows")
}

func GetMemoryUsage() (MemUsage, error) {
	return MemUsage{}, errors.New("Not implemented on Windows")
}
