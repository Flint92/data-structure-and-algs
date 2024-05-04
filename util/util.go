package util

// MaxInt return the max value both a and b
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt return the min value both a and b
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Contains[T comparable](s []T, value T) bool {
	for _, item := range s {
		if item == value {
			return true
		}
	}
	return false
}
