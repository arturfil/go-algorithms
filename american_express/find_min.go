package americanexpress

func FindMin(A []int) int {
    var result int = 0
    for i := 1; i < len(A); i++ {
        if result > A[i] {
            result = A[i]
        }
    }
    return result
}
