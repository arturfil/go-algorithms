package arrays

func maxProduct(nums []int) int {
	prod, min, max := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		tempMax := maxVal(nums[i], max*nums[i], min*nums[i])
		min = minVal(nums[i], max*nums[i], min*nums[i])
		max = tempMax
		prod = maxVal(max, prod)
	}

	return prod
}

func minVal(vals ...int) int {
	minv := vals[0]

	for _, val := range vals {
		if val < minv {
			minv = val
		}
	}

	return minv
}

func maxVal(vals ...int) int {
	maxv := vals[0]

	for _, val := range vals {
		if val > maxv {
			maxv = val
		}
	}

	return maxv
}
