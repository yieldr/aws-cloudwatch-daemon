package main

import (
	"math"
	"syscall"
)

type DiskUsageInfo struct {
	Utilization      float64
	Used             int
	Available        int
	InodeUtilization float64
}

func NewDiskUsage(path string) (*DiskUsageInfo, error) {
	s := syscall.Statfs_t{}
	err := syscall.Statfs(path, &s)
	if err != nil {
		return nil, err
	}

	total := int(s.Bsize) * int(s.Blocks)
	available := int(s.Bsize) * int(s.Bavail)

	info := &DiskUsageInfo{}
	info.Used = total - available
	info.Available = available
	info.Utilization = math.Ceil((float64(info.Used) / float64(total)) * 100)
	info.InodeUtilization = 100 * (1 - float64(s.Ffree)/float64(int(s.Files)))

	return info, nil
}
