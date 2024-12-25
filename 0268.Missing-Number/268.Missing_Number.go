func missingNumber(nums []int) int {
	num_size := len(nums)
	sum := (num_size * (num_size + 1)) / 2
	var real_sum = 0
	for i, n := range nums {
		real_sum += nums[i]
	}

	return sum - real_sum
}