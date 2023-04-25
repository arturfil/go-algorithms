package arrays

func TwoSum(nums []int, target int) []int {

	res := []int{0, 0}
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if _, ok := m[target-nums[i]]; ok {

			res = []int{m[target-nums[i]], i}
			return res

		} else {
			m[nums[i]] = i
		}

	}

	return res
}
