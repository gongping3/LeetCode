func majorityElement(nums []int) int {
	var candidate = -1
	var count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else if count--; count < 0 { // 合并减少count后判断
			candidate = nums[i]
			count = 1
		}
	}
	return candidate
}