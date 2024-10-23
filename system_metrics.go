package gometrics

type CPUUsage struct {
	User uint64
	Sys  uint64
	Idle uint64
	Nice uint64
	Wait uint64
}

type LoadAverage struct {
	One     float64
	Five    float64
	Fifteen float64
}

type MemUsage struct {
	Total uint64
	Free  uint64
	Used  uint64
}
