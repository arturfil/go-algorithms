package matrix

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	res := []int{}
	rows, cols := len(matrix), len(matrix[0])
	upper, lower, left, right := 0, rows-1, 0, cols-1

	for len(res) < rows * cols {
		for i := left; i <= right; i++ {
			res = append(res, matrix[upper][i])
		}

		for j := upper + 1; j <= lower; j++ {
			res = append(res, matrix[j][right])
		}

		if upper == lower || left == right { break }

		for i := right-1 ; i >= left; i-- {
			res = append(res, matrix[lower][i])
		}

		for j := lower-1; j >= upper+1; j-- {
			res = append(res, matrix[j][left])
		}

		upper++ 
		lower-- 
		left++ 
		right--
	}

	return res
}
