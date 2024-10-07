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
	for _, eachWord := range words {
		switch eachWord {
		case "a", "A":

		}
	}

}

/*


	switch s {
		case ".", ",", "!", "?", ";", ":" :

	}


*/
