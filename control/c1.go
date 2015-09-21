package control

// AgeCheck determines drinking condition
func AgeCheck(age int) string {
	if age < 21 {
		return `good`
	}

	return "too young"
}

// Language explores if initializer
func Language(x int) string {
	if z := x * 3; z < 30 {
		return "low"
	}

	return `hi`
}

// Looping will replicate a string
func Looping(x int, s string) string {
	for i := 0; i < x; i++ {
		s += `-`
	}

	return s
}

// Switching switches
func Switching(s string) string {
	switch s {
	case `apple`:
		return "yes"
	case `android`:
		return `maybe`
	default:
		return "nope"
	}
}
