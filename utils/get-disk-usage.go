package utils

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func GetDiskUsage(path string) string {
	var stat unix.Statfs_t
	unix.Statfs(path, &stat)

	avail := FormatBytes(int(stat.Bavail * uint64(stat.Bsize)))
	total := FormatBytes(int(stat.Blocks * uint64(stat.Bsize)))

	return fmt.Sprintf("%s of %s left", avail, total)
}
