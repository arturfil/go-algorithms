package strings

func lengthOfLongestSubstring(s string) int {
    longest, start := 0, 0
    seen := map[rune]int {}

    for i, ch := range s {
        if _, ok := seen[ch]; ok && seen[ch] >= start { // as it has not been seen
            start = seen[ch] + 1
        }

        longest = max(i - start + 1, longest)
        seen[ch] = i
    }
    return longest
}

func max(values ...int) int {
    maxVal := values[0]
    for _, val := range values {
        if val > maxVal {
            maxVal = val
        }
    }
    return maxVal
}
