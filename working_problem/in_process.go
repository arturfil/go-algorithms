package workingproblem

func isValid(s string) bool {
	bMap := map[string]string{
		{"(", ")"},
		{"[", "]"},
		{"{", "}"},
	}
	stck := []ch{}

	for _, ch := range s {
		if ch == stck[len(stck)-1] {
            stck = stck[:len(stck)-1]
		} else {
            stck = append(stck, map[ch])
		}
	}

	if len(stck) == 0 {
		return true
	} else {
		return false
	}
}
