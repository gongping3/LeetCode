func findNext(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)
	j := 0

	for i := 1; i < n; i++ {
		// 如果字符不匹配，回退到前一个可能匹配的前缀
		// 回退的思想也是利用了前缀表(即next数组)的思想，对于aaab串，aaa前缀
		// 和aab后缀匹配失败，aaa就等于是匹配串，aab就是原串，所以找aaa的前缀表
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}

		// 如果字符匹配，前缀长度加 1
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}

	return next
}

func strStr(haystack string, needle string) int {
	haystack_len := len(haystack)
	needle_len := len(needle)
	j := 0
	next := findNext(needle)

	for i := 0; i < haystack_len; i++ {
		// 如果字符不匹配，回退到前一个可能匹配的前缀
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		// 如果字符匹配
		if haystack[i] == needle[j] {
			if j == needle_len-1 {
				return i + 1 - needle_len
			}
			j++
		}
	}

	return -1
}

