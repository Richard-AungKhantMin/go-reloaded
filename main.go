package main

import (
	"os"
)

func main() {
	fileB, err := os.ReadFile("./sample.txt")
	IsErrNil(err)

	ModifiedText := MyLab(string(fileB))
	os.WriteFile("./result.txt", []byte(ModifiedText), 0644)
}
