package inprogress

func MaxSubArray(nums []int) int {
	tot, current := 0, 0

	for i := range nums {
        current = Max(nums[i], nums[i] + current)
        tot = Max(current, tot)
	}

	return tot
}

func Max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
