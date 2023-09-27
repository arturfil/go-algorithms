package twopointers

func MaxArea(height []int) int {
    max_area, left, right := 0, 0, len(height)-1

    for left < right {
        temp_h := min(height[left], height[right])
        temp_a := temp_h * (right - left)
        max_area = max(max_area, temp_a)

        if height[left] <= height[right] {
            left++
        } else {
            right--
        }
    }
    return max_area
}
