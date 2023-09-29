package arrays

func LongestConsecutive(nums []int) int {
    max_len := 0 
    set := map[int]bool{}

    for _, num := range nums { set[num] = true }

    for num := range set {
        current := num
        sum := 0
        // if key doesn't exist!
        if !set[num-1] {
            for set[current]  {
                sum++
                current++
            }
        } 
        max_len = max(max_len, sum)
    }
    return max_len
}

