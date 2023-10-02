package utils

func Max(vals ...int) int {
    maxVal := vals[0]
    for _, val := range vals {
        if val < maxVal {
            maxVal = val
        }
    }

    return maxVal
}
