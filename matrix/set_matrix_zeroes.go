package matrix

func setZeroes(matrix [][]int) {
   var colZero int = 1
    for i := 0; i < len(matrix); i++ {
        if matrix[i][0] == 0 {
            colZero = 0
        }
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    for i := len(matrix)-1; i >= 0; i-- {
        for j := len(matrix[0])-1; j > 0; j-- {
            if matrix[i][0]  == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
        if colZero == 0 { matrix[i][0] = 0 }
    } 
}
