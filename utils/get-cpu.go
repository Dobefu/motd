package utils

import (
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCpu() string {
	cpu_info, err := cpu.Info()

	if err != nil || len(cpu_info) < 1 {
		return "Unknown CPU"
	}

	return cpu_info[0].ModelName
}
