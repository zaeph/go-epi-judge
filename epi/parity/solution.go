package parity

func Parity(x int64) int16 {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	p := int16(x & 0x1)
	return p
}
