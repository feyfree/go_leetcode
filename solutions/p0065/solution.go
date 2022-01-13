package p0065

func isNumber(s string) bool {
	state := [...]map[string]int{
		{},
		{"blank": 1, "sign": 2, "digit": 3, ".": 4},
		{"digit": 3, ".": 4},
		{"digit": 3, ".": 5, "e": 6, "E": 6, "blank": 9},
		{"digit": 5},
		{"digit": 5, "e": 6, "E": 6, "blank": 9},
		{"sign": 7, "digit": 8},
		{"digit": 8},
		{"digit": 8, "blank": 9},
		{"blank": 9}}
	currentState := 1
	var cType string
	for i := 0; i < len(s); i++ {
		c := s[i]
		cType = string(c)
		if c >= '0' && c <= '9' {
			cType = "digit"
		}
		if c == ' ' {
			cType = "blank"
		}
		if c == '+' || c == '-' {
			cType = "sign"
		}
		if _, ok := state[currentState][cType]; !ok {
			return false
		}
		currentState = state[currentState][cType]
	}
	if currentState != 3 && currentState != 5 && currentState != 8 && currentState != 9 {
		return false
	}
	return true
}
