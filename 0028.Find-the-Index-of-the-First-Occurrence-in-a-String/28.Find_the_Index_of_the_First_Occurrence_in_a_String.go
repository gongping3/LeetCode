func strStr(haystack string, needle string) int {
	haystack_len := len(haystack)
	needle_len := len(needle)
	j := 0
	for i := 0; i < haystack_len; i++ {
		if haystack[i] == needle[j] {
			if j == needle_len-1 {
				return i + 1 - needle_len
			}
			j++
		} else {
			i = i - j
			j = 0
		}
	}
	return -1
}