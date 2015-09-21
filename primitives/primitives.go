package primitives

// Sum adds two numbers
func Sum(x, y int) int {
	return x + y
}

// Double will 2x the input
func Double(x float64) int64 {
	y := int64(x)
	y <<= 1
	return y
}

// Concat will add the two strings
func Concat(s1, s2 string) string {
	return s1 + s2
}
