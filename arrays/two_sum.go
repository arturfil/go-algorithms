package arrays

func TwoSum(nums []int, target int) []int {
	numsMap := map[int]int{}

	for i, num := range nums {
		if _, exists := numsMap[num]; exists {
			return []int{numsMap[num], i}
		} else {
			numsMap[target-num] = i
		}
	}

	return []int{0, 0}
}
