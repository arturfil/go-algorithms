package strings

import "fmt"

func GroupAnagrams(strs []string) [][]string {
    var arr [][]string
    for i := 0; i < len(strs); i++ {
        fmt.Printf("%v\n", strs[i])
    }
    return arr
}
