
function exist(board: string[][], word: string): boolean {
    for(let i = 0; i < board.length; i++) {
        for(let j = 0; j < board[i].length; j++) {
            if (dfs(board,i ,j , word, 0)) return true;
        }
    }
    return false;
};

function dfs(board:string[][], row:number, col:number, word:string, idx:number) {
    if (word.length === idx) return true;

    if (
        row >= board.length || 
        row < 0 || 
        board[row][col] !== word[idx]
    ) return false;

    board[row][col] = "#"
    
    if (
        dfs(board, row + 1, col, word, idx + 1) ||
        dfs(board, row - 1, col, word, idx + 1) ||
        dfs(board, row, col + 1, word, idx + 1) ||
        dfs(board, row, col - 1, word, idx + 1)
    ) return true; // this will be true if word.length === idx;

    board[row][col] = word[idx]; // reset.
}
