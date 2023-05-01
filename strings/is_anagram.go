package strings

func IsAnagram(t string, s string) bool {
    if (len(s) != len(t)) { return false }
    var arr [26]int

    for _, ch := range(t) {
        idx := int(ch - rune('a')  )
        arr[idx] += 1
    }

    for _, ch := range(s) {
        idx := int(ch - rune('a')  )
        if arr[idx] == 0 { return false }
        arr[idx] -= 1
    }
    return true
}



