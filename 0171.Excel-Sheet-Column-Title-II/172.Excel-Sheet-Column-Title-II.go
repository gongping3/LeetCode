import (
	"math"
)

func titleToNumber(columnTitle string) int {
	col_len := len(columnTitle)
	i := 0
	sum := 0
	for col_len-i > 0 {
		val := int(columnTitle[col_len-i-1]) - int('A') + 1
		if i == 0 {
			sum = sum + val
		} else {
			sum = sum + val*int(math.Pow(float64(26), float64(i)))
		}
		i++
	}
	return sum
}