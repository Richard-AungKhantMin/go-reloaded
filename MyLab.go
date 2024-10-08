package main

import (
	"strings"
)

func MyLab(sentence string) string {

	words := strings.Fields(sentence)

	if len(words) > 0 {
		words[0] = CapCap(words[0])
	}

	var QuoteCount int
	var num int

	for i, eachWord := range words {
		switch eachWord {

		case "a", "A":
			if strings.Contains("aeiouhAEIOUH", string(words[i][0])) {
				eachWord = eachWord + "n"
			}

		case "(hex)":
			words[i-1] = HexToDec(words[i-1])
			eachWord = ""

		case "(bin)":
			words[i-1] = BinaryToDec(words[i-1])
			eachWord = ""

		case "(up)":
			words[i-1] = Upper(words[i-1])
			eachWord = ""

		case "(low)":
			words[i-1] = Lower(words[i-1])
			eachWord = ""

		case "(cap)":
			words[i-1] = CapCap(words[i-1])
			eachWord = ""

		case "(low,":
			num = NumExtractor(words[i+1])
			NumMod(num, i, words, Lower)
			eachWord = ""

		case "(up":
			num = NumExtractor(words[i+1])
			NumMod(num, i, words, Upper)
			eachWord = ""

		case "(cap":
			num = NumExtractor(words[i+1])
			NumMod(num, i, words, CapCap)
			eachWord = ""

		case "'":
			if QuoteCount%2 == 1 {
				words[i+1] = "'" + words[i+1]
				eachWord = ""
				QuoteCount++
			} else {
				words[i-1] = words[i+1] + "'"
				eachWord = ""
			}

		default:
			for strings.Contains(".,!?;:", string(eachWord[0])) && i != 0 {
				words[i-1] = words[i-1] + string(eachWord[0])
				eachWord = eachWord[1:]
			}

		}
	}

	return strings.Join(words, " ")
}

/*


	switch s {
		case ".", ",", "!", "?", ";", ":" :

	}


*/
