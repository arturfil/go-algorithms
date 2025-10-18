package main

import "leetcode/matrix"

// import "leetcode/helpers"

func main() {
	// helpers.ChooseRandomProblem()
	// workingproblem.ProductExceptSelf([]int{1,2,3,4}) // 24, 12, 8, 6

	/*
		[
			1, 2, 3
			4, 5, 6
			7, 8, 9
		]

	*/

	test := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	matrix.SpiralOrder(test)

}
