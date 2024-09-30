package functions

import (
	"strconv"
)

func BinaryToDec(input string) string {

	BiValue, err := strconv.ParseInt(input, 2, 64)
	IsErrNil(err)

	return strconv.Itoa(int(BiValue))
}
