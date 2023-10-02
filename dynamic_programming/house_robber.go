package dynamicprogramming

func Rob(nums []int) int {
    prevMax, currMax := 0, 0 

    for _, house := range nums {
        temp := max(house + prevMax, currMax)
        prevMax = currMax
        currMax = temp
    }

    return currMax
}
