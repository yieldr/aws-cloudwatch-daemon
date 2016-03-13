package main

import "testing"

type convFn func(uint64) uint64

func TestBytes(t *testing.T) {
	for b, xb := range map[uint64]uint64{
		1024:          1, // KB
		1050:          1,
		2048:          2,
		2050:          2,
		1048576:       1, // MB
		1050000:       1,
		2097152:       2,
		2100000:       2,
		1073741824:    1, // GB
		1100000000:    1,
		2147483648:    2,
		2150000000:    2,
		1099511627776: 1, // TB
		1100000000000: 1,
		2199023255552: 2,
		2200000000000: 2,
	} {
		var fn convFn
		switch {
		case b >= TB:
			fn = ToTerraBytes
		case b >= GB:
			fn = ToGigaBytes
		case b >= MB:
			fn = ToMegaBytes
		case b >= KB:
			fn = ToKiloBytes
		}
		if xb != fn(b) {
			t.Errorf("failed asserting that %dB equals %dxB, got %dxB instead", b, xb, fn(b))
		}
	}
}
