package arrays

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{0, 0}
}
