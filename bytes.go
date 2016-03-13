package main

const (
	B  = 1.0
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

func ToKiloBytes(bytes uint64) uint64 {
	return bytes / KB
}

func ToMegaBytes(bytes uint64) uint64 {
	return bytes / MB
}

func ToGigaBytes(bytes uint64) uint64 {
	return bytes / GB
}

func ToTerraBytes(bytes uint64) uint64 {
	return bytes / TB
}
