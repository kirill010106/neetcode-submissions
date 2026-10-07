import "slices"

func hasDuplicate(nums []int) bool {
    // m := make(map[int]struct{})
    // for _, v := range nums {
    //     if _, ok := m[v]; ok {
    //         return true
    //     }
    //     m[v] = struct{}{}
    // }
    // return false
    slices.Sort(nums)
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            return true
        }
    }
    return false
}
