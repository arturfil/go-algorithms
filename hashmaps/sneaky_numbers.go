package hashmaps

func GetSneakyNumbers(nums []int) []int {
    repeated := map[int]int{}    
    sneaky := []int{}

    for _, num := range nums {
        if _, exists := repeated[num]; !exists {
            repeated[num] = 1
        } else {
            repeated[num] += 1
        }
    }

    for k, v := range repeated {
        if v == 2 {
            sneaky = append(sneaky, k)
        }
    }

    return sneaky
}

