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
	
	it := b.Itor()
	for {
		ch, idx, err := it.Next()
		if err != nil {
			break
		}
		if idx%8 == 7 {
			print(ch, "_")
		} else {
			print(ch)
		}
	}
}

func printByteArray(bts []byte) {
	fmt.Print("[")
	for _, v := range bts {
		fmt.Printf("%x,", v)
	}
	fmt.Print("]\n")
}
