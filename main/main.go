package main

import "leetcode/graphs"

// "leetcode/helpers"


func main() {
    board := [][]byte {
        {'#', 'u', '#', 'l', 'o'},
        {'y', 'o', 'u', 's', 'e'},
        {'t', 'u', 't', 'b', 'e'},
    }

    word := "usb"

    /*
    step1 - go through each character until I find a matching one
    step2 - dfs onto a neighbor one
    step3 - make sure that I can get the full word path
        if get the full path 
            return True
        else 
            return False

    dfs:
        

    */


    // helpers.ChooseRandomProblem()
    graphs.Exist(board, word)
}
