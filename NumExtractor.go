package main

import "strconv"

func NumExtractor(s string) int {
	Num, err := strconv.Atoi(string(s[:len(s)-1]))
	IsErrNil(err)
	return Num
}
