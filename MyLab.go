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
	var num

	for i, eachWord := range words {
		switch eachWord {

		case "a", "A":
			if strings.Contains("aeiouhAEIOUH", string(words[i][0])){
				eachWord = eachWord + "n"
			}

		case "(hex)":
			words[i-1] = HexToDec(words[i-1])

		case "(bin)":
			words[i-1] = BinaryToDec(words[i-1])

		case "(up)":
			words[i-1] = Upper(words[i-1])

		case "(low)":
			words[i-1] = Lower(words[i-1])

		case "(cap)":
			words[i-1] = CapCap(words[i-1])

		case "(low,":


		case "(up":

		case "(cap":

		case "'":

		default:

		}
	}


}

/*


	switch s {
		case ".", ",", "!", "?", ";", ":" :

	}


*/
