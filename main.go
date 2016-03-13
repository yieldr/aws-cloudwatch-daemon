package main

import (
	"flag"
	"log"
	"os"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var args struct {
	memoryUsage bool
	diskUsage   bool
	diskPath    string
	namespace   string
}

func init() {
	flag.BoolVar(&args.memoryUsage, "memory-usage", true, "Collect memory statistics (total, used, available, utilization, swap used, swap utilization)")
	flag.BoolVar(&args.diskUsage, "disk-usage", true, "Collect disk space statistics (total, used, available, utilization, inode utilization)")
	flag.StringVar(&args.diskPath, "disk-path", "/", "Disk Path")
	flag.StringVar(&args.namespace, "aws-cloudwatch-ns", "CoreOS", "CloudWatch metric namespace (required)")
	flag.Parse()
}

func main() {
	metadata, err := NewEc2Metadata()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	cw := NewCloudWatch(args.namespace, metadata.Region)
	cw.AddDimention("InstanceId", metadata.InstanceID)
	cw.AddDimention("InstanceType", metadata.InstanceType)
	cw.AddDimention("ImageId", metadata.ImageID)

	if args.diskUsage {
		du, err := disk.DiskUsage(args.diskPath)
		if err != nil {
			log.Printf("Failed to get disk usage: %s", err)
			os.Exit(2)
		}
		cw.AddMetric("DiskTotal", "Gigabytes", float64(ToGigaBytes(du.Total)))
		cw.AddMetric("DiskUsed", "Gigabytes", float64(ToGigaBytes(du.Used)))
		cw.AddMetric("DiskFree", "Gigabytes", float64(ToGigaBytes(du.Free)))
		cw.AddMetric("DiskUsedPercent", "Percent", du.UsedPercent)

		cw.AddMetric("DiskInodesTotal", "Gigabytes", float64(ToGigaBytes(du.InodesTotal)))
		cw.AddMetric("DiskInodesUsed", "Gigabytes", float64(ToGigaBytes(du.InodesUsed)))
		cw.AddMetric("DiskInodesFree", "Gigabytes", float64(ToGigaBytes(du.InodesFree)))
		cw.AddMetric("DiskInodesUsedPercent", "Percent", du.InodesUsedPercent)
	}

	if args.memoryUsage {
		vm, err := mem.VirtualMemory()
		if err != nil {
			log.Printf("Failed to get memory usage: %s", err)
			os.Exit(2)
		}
		cw.AddMetric("MemoryTotal", "Gigabytes", float64(ToGigaBytes(vm.Total)))
		cw.AddMetric("MemoryUsed", "Gigabytes", float64(ToGigaBytes(vm.Used)))
		cw.AddMetric("MemoryFree", "Gigabytes", float64(ToGigaBytes(vm.Free)))
		cw.AddMetric("MemoryUsedPercent", "Percent", vm.UsedPercent)

		sw, err := mem.SwapMemory()
		if err != nil {
			log.Printf("Failed to get swap memory usage: %s", err)
			os.Exit(2)
		}
		cw.AddMetric("SwapTotal", "Gigabytes", float64(ToGigaBytes(sw.Total)))
		cw.AddMetric("SwapUsed", "Gigabytes", float64(ToGigaBytes(sw.Used)))
		cw.AddMetric("SwapFree", "Gigabytes", float64(ToGigaBytes(sw.Free)))
		cw.AddMetric("SwapUsedPercent", "Percent", sw.UsedPercent)
	}

	_, err = cw.Send()
	if err != nil {
		log.Printf("Failed to send metrics to cloudfront: %s", err)
		os.Exit(3)
	}

	log.Printf("Successfully sent metrics to AWS CloudWatch.")
}
