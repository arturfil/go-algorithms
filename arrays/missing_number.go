package arrays

func missingNumber(nums []int) int {
    org, sum := 0, 0

    for _, num := range nums {
        sum += num
    }

    for i := 0; i < len(nums)+1; i++ {
        org += i
    } 

    return org - sum

}
