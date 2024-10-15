package main

import (
	"strings"
)

func MyLab(sentence string) string {
	words := strings.Fields(sentence)
	if len(words) > 0 {
		words[0] = CapCap(words[0])
	}
	var num int
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			if i > 0 {
				words[i-1] = HexToDec(words[i-1])
				words[i] = ""
			}
		case "(bin)":
			if i > 0 {
				words[i-1] = BinaryToDec(words[i-1])
				words[i] = ""
			}
		case "(up)":
			if i > 0 {
				words[i-1] = Upper(words[i-1])
				words[i] = ""
			}
		case "(low)":
			if i > 0 {
				words[i-1] = Lower(words[i-1])
				words[i] = ""
			}
		case "(cap)":
			if i > 0 {
				words[i-1] = CapCap(words[i-1])
				words[i] = ""
			}
		case "(low,":
			if i < len(words)-1 {
				num = NumExtractor(words[i+1])
				NumMod(num, i, words, Lower)
				words[i] = ""
				words[i+1] = ""
			}
		case "(up,":
			if i < len(words)-1 {
				num = NumExtractor(words[i+1])
				NumMod(num, i, words, Upper)
				words[i] = ""
				words[i+1] = ""
			}
		case "(cap,":
			if i < len(words)-1 {
				num = NumExtractor(words[i+1])
				NumMod(num, i, words, CapCap)
				words[i] = ""
				words[i+1] = ""
			}
		}
	}
	edited := Vowels(QuoteHandler(PunPun(SliceSlayer(words))))
	return strings.Join(edited, " ")
}
