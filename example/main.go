package main

import (
	"fmt"

	"github.com/shizuku/gobits"
)

func main() {
	x0 := gobits.FromByte(0, 1)
	println(x0.String())
	x10 := gobits.FromByte(2, 2)
	println(x10.String())
	x110 := gobits.FromByte(6, 3)
	println(x110.String())
	x1110 := gobits.FromByte(14, 4)
	println(x1110.String())
}

func printByteArray(bts []byte) {
	fmt.Print("[")
	for _, v := range bts {
		fmt.Printf("%x,", v)
	}
	fmt.Print("]\n")
}
