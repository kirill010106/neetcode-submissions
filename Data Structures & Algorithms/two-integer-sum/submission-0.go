func retarr(a int, b int) []int {
	if a > b {
		return []int{b, a}
	}
	return []int{a, b}
}

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	for i := range nums {
		if val, ok := m[nums[i]]; ok {
			return retarr(i, val)
		}
		m[target-nums[i]] = i
	}
	return []int{}
}
