package helpers

func Max(values ...int) int {
    maxVal := values[0]
    for _, val := range values {
        if val > maxVal {
            maxVal = val
        }
    }
    return maxVal
}
