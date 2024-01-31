package strings

func CountSubstrings(s string) int {
    sum, odd, even := 0, 0, 0 
    for i := 0; i < len(s); i++ {
        odd = subPalindrome(s, i, i)
        even = subPalindrome(s, i, i+1)
        sum += (odd + even)
    }
    return sum
}

func subPalindrome(s string, l, r int) int {
    count := 0
    for l >= 0 && r < len(s) && s[l] == s[r] { // different from typescript
        l--
        r++
        count++
    }
    return count
}
