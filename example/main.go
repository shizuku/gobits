package main

import (
	"fmt"

	"github.com/shizuku/gobits"
)

func main() {
	b := gobits.New()

	b.AppendBits(0xab, 5)
	println(b.String())
	b.AppendByte(0xcd)
	b.AppendBits(0xcd, 7)
	println(b.String())
}

func printByteArray(bts []byte) {
	fmt.Print("[")
	for _, v := range bts {
		fmt.Printf("%x,", v)
	}
	fmt.Print("]\n")
}
