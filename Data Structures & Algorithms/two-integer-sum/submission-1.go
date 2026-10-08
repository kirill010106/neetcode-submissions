func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	for i := range nums {
		if val, ok := m[target-nums[i]]; ok {
			return []int{val, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
