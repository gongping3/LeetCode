func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func moveZeroes(nums []int) {
	left := 0
	right := 0
	num_size := len(nums)
	for right < num_size {
		if nums[right] != 0 {
			swap(&nums[right], &nums[left])
			left++
		}
		right++
	}
}