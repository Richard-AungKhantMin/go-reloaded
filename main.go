package main

import (
	"fmt"
	"os"
)

func main() {
	fileB, err := os.ReadFile("./sample.txt")
	IsErrNil(err)

	ModifiedText := MyLab(string(fileB))
	fmt.Println(ModifiedText)
	os.WriteFile("./result.txt", []byte(ModifiedText), 0644)
}
