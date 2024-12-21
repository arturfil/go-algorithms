package arrays

func containsDuplicate(nums []int) bool {
    set := map[int]bool{}
    
    for _, num := range nums {
        if _, has := set[num]; has {
            return true
        } else {
            set[num] = true
        }
    }

    return false
}
