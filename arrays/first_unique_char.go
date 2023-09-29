package arrays

func FirstUniqChar(s string) int {
    prev := map[rune]int {} 

    // place or add a 1 on the value portion of the map
    for _, ch := range s {
        prev[ch]++
    }

    // it's only seen once, thus unique, return i
    for i, ch := range s {
        if prev[ch] == 1 { 
           return i 
        }
    }

    return -1
}
