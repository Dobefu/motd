package utils

import (
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCpu() string {
	cpuInfo, err := cpu.Info()

	if err != nil || len(cpuInfo) < 1 {
		return "Unknown CPU"
	}

	return cpuInfo[0].ModelName
}
