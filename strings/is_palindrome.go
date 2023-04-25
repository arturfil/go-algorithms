package strings

import "fmt"

func IsPalindrome(s string) bool {
	fmt.Println("test")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ,", s[i])
	}
	return false
}
