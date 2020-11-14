package main

import (
	"fmt"

	"github.com/shizuku/gobits"
)

func main() {
	b := gobits.New()

	b.AppendBit(1)
	println(b.String())
	b.AppendBit(0)
	println(b.String())
	b.AppendBit(0)
	println(b.String())
	b.AppendBit(1)
	println(b.String())
	b.AppendBit(1)
	println(b.String())
	b.AppendBits(0xab, 8)
	println(b.String())
	printByteArray(b.Bytes())
	a := gobits.FromBytes(b.Bytes(), 0)
	println(a.String())
}

func printByteArray(bts []byte) {
	fmt.Print("[")
	for _, v := range bts {
		fmt.Printf("%x,", v)
	}
	fmt.Print("]\n")
}
