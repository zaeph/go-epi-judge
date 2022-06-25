package parity

// func Parity(x int64) int16 {
// 	// Bruteforce
// 	var p int64
// 	p = 0
// 	for x > 0 {
// 		p ^= x & 0x1
// 		x >>= 1
// 	}
// 	return int16(p)
// }

// func Parity(x int64) int16 {
// 	// Optimized
// 	var p int16
// 	p = 0
// 	for x > 0 {
// 		p ^= 1
// 		x = x & (x - 1)
// 	}
// 	return p
// }

func Parity(x int64) int16 {
	// Best
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return int16(x & 0x1)
}
