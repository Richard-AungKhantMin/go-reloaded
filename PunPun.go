package main

import "strings"

func PunPun(words []string) []string {

	for i := 0; i < len(words); i++ {
		for i > 0 && len(words[i]) > 0 && strings.Contains(".,!?;:", string(words[i][0])) {
			words[i-1] = words[i-1] + string(words[i][0])
			words[i] = words[i][1:]
		}
	}
	return SliceSlayer(words)
}
