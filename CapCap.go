package main

import "unicode"

func CapCap(input string) string {

	slicy := []rune(input)
	var answer string

	slicy[0] = unicode.ToUpper(slicy[0])

	if len(slicy) > 1 {
		for i := 1; i < len(slicy); i++ {
			slicy[i] = unicode.ToLower(slicy[i])
		}
	}

	for _, each := range slicy {
		answer = answer + string(each)
	}

	return answer
}
