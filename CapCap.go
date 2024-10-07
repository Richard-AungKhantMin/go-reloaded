package main

func CapCap(input string) string {

	if len(input) == 1 {
		return Upper(string(input))
	}

	return Upper(string(input[0])) + Lower(string(input[1:]))
}
