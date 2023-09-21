package arrays

func ProductsExceptSelf(nums []int) []int {
    var res = make([]int, len(nums))
    
    add := 1
    for i := range nums {
        res[i] = add
        add *= nums[i]
    }

    add = 1
    for i := len(nums)-1; i >= 0; i-- {
        res[i] = add * res[i]
        add *= nums[i]
    } 

    return res
}
