package tree

// maxInt return the max value both a and b
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// minInt return the min value both a and b
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func contains[T comparable](s []T, value T) bool {
	for _, item := range s {
		if item == value {
			return true
		}
	}
	return false
}
