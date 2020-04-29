package lengthstring

func lengthOfLongestSubstring(s string) int {
	current := ""
	max := 0
	m := make(map[string]bool)

	for i := 0; i < len(s); i++ {
		current += string(s[i])
		if _, ok := m[string(s[i])]; ok {
			for {
				if current[0] == s[i] {
					current = current[1:]
					break
				} else {
					delete(m, string(current[0]))
					current = current[1:]
				}
			}
		} else {
			m[string(s[i])] = true
		}
		if len(current) > max {
			max = len(current)
		}

	}
	return max
}
