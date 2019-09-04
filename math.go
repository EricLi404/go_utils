package go_utils

//AbsInt64 returns the absolute value of int64
func AbsInt64(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
