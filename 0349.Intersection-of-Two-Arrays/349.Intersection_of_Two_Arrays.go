func intersection(nums1 []int, nums2 []int) []int {
	intersection := make(map[int]bool)
	ret := []int{}
	for i := 0; i < len(nums1); i++ {
		_, ok := intersection[nums1[i]]
		if ok == false {
			intersection[nums1[i]] = true
		}
	}
	for j := 0; j < len(nums2); j++ {
		if _, ok := intersection[nums2[j]]; ok {
			ret = append(ret, nums2[j])
			delete(intersection, nums2[j])
		}
	}
	return ret
}