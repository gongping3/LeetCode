func getInt(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return +50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	if len(s) < 1 {
		return 0
	}

	// 定义罗马数字的对应值
	// romanMap := map[byte]int{
	//     'I': 1,
	//     'V': 5,
	//     'X': 10,
	//     'L': 50,
	//     'C': 100,
	//     'D': 500,
	//     'M': 1000,
	// }

	sum := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- { // 从右向左遍历
		// currentValue := romanMap[s[i]]
		currentValue := getInt(s[i])

		if currentValue < prevValue { // 如果当前值小于之前的值，则执行减法
			sum -= currentValue
		} else { // 否则执行加法
			sum += currentValue
		}

		prevValue = currentValue
	}

	return sum
}
