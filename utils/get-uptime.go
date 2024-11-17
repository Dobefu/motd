package utils

import (
	"github.com/shirou/gopsutil/v4/host"
)

func GetUptime() string {
	info, err := host.Info()

	if err != nil {
		return "Uptime unknown"
	}

	uptime := info.Uptime

	return Duration(float64(uptime))
}
