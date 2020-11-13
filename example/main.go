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
	b.AppendBits(0xac, 4)
	println(b.String())

}

func printByteArray(bts []byte) {
	fmt.Print("[")
	for _, v := range bts {
		fmt.Printf("%x,", v)
	}
	fmt.Print("]\n")
}
