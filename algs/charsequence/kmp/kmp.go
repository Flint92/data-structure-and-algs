package kmp

func KMP(s, p string) []int {
	rs := []rune(s)
	rp := []rune(p)

	i := 0
	j := 0
	nxt := next(rp)

	result := make([]int, 0)
	for i < len(rs) && j < len(rp) {
		if j == -1 || rs[i] == rp[j] {
			i++
			j++
			if j == len(rp) {
				result = append(result, i-j)
				j = 0
			}
		} else {
			j = nxt[j]
		}
	}

	return result
}

func next(p []rune) []int {
	nxt := make([]int, len(p))
	nxt[0] = -1

	i := 0
	j := -1
	for i < len(p)-1 {
		if j == -1 || p[i] == p[j] {
			i++
			j++
			if p[i] == p[j] {
				nxt[i] = nxt[j]
			} else {
				nxt[i] = j
			}

		} else {
			j = nxt[j]
		}
	}

	return nxt
}
