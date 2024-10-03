package main

import (
	"os"
)

func main() {
	file, err := os.ReadFile("sample.txt")
	IsErrNil(err)

	fileS := string(file)

}
