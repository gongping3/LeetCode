func intersect(nums1 []int, nums2 []int) []int {
	intersection := make(map[int]int)
	ret := []int{}
	for i := 0; i < len(nums1); i++ {
		intersection[nums1[i]]++

	}
	for j := 0; j < len(nums2); j++ {
		if _, ok := intersection[nums2[j]]; ok {
			ret = append(ret, nums2[j])
			intersection[nums2[j]]--
			if intersection[nums2[j]] == 0 {
				delete(intersection, nums2[j])
			}
		}
	}
	return ret
}