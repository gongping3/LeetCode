import "fmt"

func isAlphanumeric(b byte) bool {
	return (b >= 48 && b <= 57) || (b >= 65 && b <= 90) || (b >= 97 && b <= 122)
}

func isPalindrome(s string) bool {
	var slice []byte
	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			slice = append(slice, s[i])
		}
	}
	slice_len := len(slice)
	i := 0
	j := slice_len - 1
	for i < j {
		if slice[i] != slice[j] && (slice[i]^32) != slice[j] {
			return false
		}
		i++
		j--
	}
	return true
}