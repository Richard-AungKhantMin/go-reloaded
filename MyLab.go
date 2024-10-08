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

	for i := 0; i < len(words); i++ {
		switch words[i] {

		case "a", "A":
			if strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) {
				words[i] = words[i] + "n"
			}

		case "(hex)":
			words[i-1] = HexToDec(words[i-1])
			words[i] = ""

		case "(bin)":
			words[i-1] = BinaryToDec(words[i-1])
			words[i] = ""

		case "(up)":
			words[i-1] = Upper(words[i-1])
			words[i] = ""

		case "(low)":
			words[i-1] = Lower(words[i-1])
			words[i] = ""

		case "(cap)":
			words[i-1] = CapCap(words[i-1])
			words[i] = ""

		case "(low,":
			num = NumExtractor(words[i+1])
			NumMod(num, i, words, Lower)
			words[i] = ""
			words[i+1] = ""

		case "(up,":
			num = NumExtractor(words[i+1])
			NumMod(num, i, words, Upper)
			words[i] = ""
			words[i+1] = ""

		case "(cap,":
			num = NumExtractor(words[i+1])
			NumMod(num, i, words, CapCap)
			words[i] = ""
			words[i+1] = ""

		case "'":
			if QuoteCount%2 == 0 {
				words[i+1] = "'" + words[i+1]
				words[i] = ""
				QuoteCount++
			} else {
				words[i-1] = words[i-1] + "'"
				words[i] = ""
				QuoteCount++
			}
		}
	}

	edited := SliceSlayer(words)
	edited = SliceSlayer(PunPun(edited))

	return strings.Join(edited, " ")
}

/*


	switch s {
		case ".", ",", "!", "?", ";", ":" :

	}


*/
