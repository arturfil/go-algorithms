package strings

func IsValid(s string) bool {
   p := map[rune]rune{
       '{':'}',
       '(':')',
       '[':']',
   } 
    stack := []rune{} // emulate a stack

    for _, c := range s {
        if len(stack) > 0 && c == stack[len(stack)-1] {
            stack = stack[:len(stack)-1] // pop
        } else {         
            stack = append(stack, p[c])  // push
        }
   }

   return len(stack) == 0
} 
