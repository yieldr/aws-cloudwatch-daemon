package main

import (
	"flag"
	"log"
	"os"
)

var args struct {
	memory    bool
	disk      bool
	diskPath  string
	namespace string
}

func init() {
	flag.BoolVar(&args.memory, "memory-usage", true, "Collect memory statistics (utilization, used, available, swap utilization, swap used)")
	flag.BoolVar(&args.disk, "disk-usage", true, "Collect disk space statistics (utilization, used, available, inode utilization)")
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

	if args.disk {
		info, err := NewDiskUsage(args.diskPath)
		if err != nil {
			log.Printf("Failed to get disk usage: %s", err)
			os.Exit(2)
		}
		cw.AddMetric("DiskUtilization", "Percent", info.Utilization)
		cw.AddMetric("DiskUsed", "Bytes", float64(info.Used))
		cw.AddMetric("DiskAvailable", "Bytes", float64(info.Available))
		cw.AddMetric("DiskInodesUtilization", "Percent", info.InodeUtilization)
	}

	if args.memory {
		info, err := NewMemoryUsage()
		if err != nil {
			log.Printf("Failed to get memory usage: %s", err)
			os.Exit(2)
		}

		cw.AddMetric("MemoryUtilization", "Percent", info.Utilization)
		cw.AddMetric("MemoryUsed", "Bytes", info.Used)
		cw.AddMetric("MemoryAvailable", "Bytes", info.Available)
		cw.AddMetric("SwapUtilization", "Percent", info.SwapUtilization)
		cw.AddMetric("SwapUsed", "Bytes", info.SwapUsed)
	}

	_, err = cw.Send()
	if err != nil {
		log.Printf("Failed to send metrics to cloudfront: %s", err)
		os.Exit(3)
	}

	log.Printf("Successfully sent metrics to AWS CloudWatch.")
}
