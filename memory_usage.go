package main

import (
	"math"

	"github.com/guillermo/go.procmeminfo"
)

type MemoryUsageInfo struct {
	Utilization,
	Used,
	Available,
	SwapUtilization,
	SwapUsed float64
}

func NewMemoryUsage() (*MemoryUsageInfo, error) {
	meminfo := &procmeminfo.MemInfo{}
	meminfo.Update()

	return &MemoryUsageInfo{
		Utilization:     math.Ceil(float64(meminfo.Used())/float64(meminfo.Total())) * 100,
		Used:            float64(meminfo.Used()),
		Available:       float64(meminfo.Available()),
		SwapUtilization: math.Ceil(float64(meminfo.Swap())),
		SwapUsed:        float64((*meminfo)["SwapCached"]),
	}, nil
}
