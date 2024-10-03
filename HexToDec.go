package main

import (
	"strconv"
)

func HexToDec(input string) string {

	decimalValue, err := strconv.ParseInt(input, 16, 64)
	IsErrNil(err)

	return strconv.Itoa(int(decimalValue))
}
