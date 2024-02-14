package graphs

func Exist(board [][]byte, word string) bool {

	for i := 0; i < len(board); i++ { // single row length
		for j := 0; j < len(board[i]); j++ {

            if dfs(board, i, j, 0, word) {
                return true
            }

		}
	}

	return false
}

func dfs(board [][]byte, i, j, idx int, word string) bool {
	if idx == len(word) { return true } // if we reach this point, we are done

	// return false if goes past bounds or word[idx] != board[row][col]
	if  i < 0 || i >= len(board)    || 
        j < 0 || j >= len(board[0]) || 
        board[i][j] == '#'          || 
        word[idx] != board[i][j]      {
		return false
	}

	board[i][j] = '#' // visited

	// upper bound && word[idx] == board[i][j]
	if  dfs(board, i-1, j, idx+1, word) || 
        dfs(board, i+1, j, idx+1, word) ||
	    dfs(board, i, j-1, idx+1, word) || 
        dfs(board, i, j+1, idx+1, word)   {
		return true
	}
	// no neighbors and never reached base case
	board[i][j] = word[idx] // reset for other recursions that are succesful at the momemtn
	return false
}

