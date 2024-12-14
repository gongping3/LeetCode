func searchInsert(nums []int, target int) int {
    end := len(nums) - 1
    begin := 0
    mid := 0

    for (begin <= end) {
        mid = (begin + end) / 2
        if (nums[mid] == target) {
            return mid
        } else if (nums[mid] < target) {
            begin = mid + 1
        } else {
            end = mid - 1
        }
    }
    return begin
}
