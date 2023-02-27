package main

func rgbToBw(r, g, b uint32) uint32 {
	return uint32(0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b))
}
