package strings

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
    re := regexp.MustCompile("[^a-zA-Z0-9]+")
    s = strings.ToLower(re.ReplaceAllString(s, ""))
    
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] { return false }
        left++ 
        right--
    }
    
	return true 
}
