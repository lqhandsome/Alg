package main

func forceSearch(s, p string) int {

	if p == "" {
		return -1
	}

	if len(p) > len(s) {
		return -1
	}

	for i := 0; i <= len(s)-len(p); i++ {
		var j int
		for ; j < len(p); j++ {
			if s[i+j] != p[j] {
				break
			}
		}
		if j == len(p) {
			return i
		}
	}

	return -1
}
