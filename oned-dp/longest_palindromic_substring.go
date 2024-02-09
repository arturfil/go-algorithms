package oneddp

func LongestPalindrome(s string) string {
    current, max := "", ""

	for i := 0; i < len(s); i++ {
		odd := expandAround(s, i, i+1)
		even := expandAround(s, i, i)

		if len(even) > len(odd) { current = even } else { current = odd }
		if len(current) > len(max) { max = current }
	}

	return max
}

func expandAround(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l-- 
		r++ 
	}
	return s[l+1 : r]
}

