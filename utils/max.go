package utils

type Numbers interface {
    int | float32 | float64
}

// where Num is type of Numbers and thus N would be int | float...
func Max[Num Numbers](vals ...Num) Num {
    maxVal := vals[0]
    for _, val := range vals {
        if val < maxVal {
            maxVal = val
        }
    }

    return maxVal
}
