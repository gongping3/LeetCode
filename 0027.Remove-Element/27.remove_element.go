func removeElement(nums []int, val int) int {
    var new_array_index int = 0

    for old_array_index := 0; old_array_index < len(nums); old_array_index++ { 
        if nums[old_array_index] != val {
            nums[new_array_index] = nums[old_array_index]
            new_array_index++
        }
    }

    return new_array_index
}