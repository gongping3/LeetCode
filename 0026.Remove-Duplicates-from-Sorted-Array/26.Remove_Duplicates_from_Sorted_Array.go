func removeDuplicates(nums []int) int {
	var i int = 0

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
            nums[i] = nums[j]
		} 
	}
    i++

	return i
}