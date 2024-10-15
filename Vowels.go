package main

import "strings"

func Vowels(words []string) []string {
	for i := 0; i < len(words); i++ {
		if strings.Contains("Aa", words[i]) && i < len(words)-1 {
			if strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) {
				words[i] = words[i] + "n"
			}
		}
	}
	return words
}
