package main

import (
	// "leetcode/helpers"
	"fmt"
	dynamicprogramming "leetcode/dynamic_programming"
)

func main() {
    // helpers.RandomProblem()
    nums := []int {10,9,2,5,3,7,101,18}
    fmt.Println(dynamicprogramming.LengthOfLIS(nums))
    // words := []string {"one", "two", "noe", "tow"}
    // strings.GroupAnagrams(words)
    // strings.IsAnagram("racecar", "cracera")
}
