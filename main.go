package main

import (
	"fmt"
	"go-reloaded/functions"
)

func main() {

	fmt.Println(functions.HexToDec("1E"))
	fmt.Println(functions.BinaryToDec("1000"))
	fmt.Println((functions.Lower("HELLOOOOOO")))
	fmt.Println(functions.Upper("let me get out of here!"))
	fmt.Println(functions.CapCap("ålANds"))
}
