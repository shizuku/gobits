package main

import (
	"github.com/shizuku/gobits"
)

func main() {
	bits := gobits.FromByte(0xff, 8)
	bits.Print()

	// by := 0xff
	// fmt.Printf("%#b\n", (by & (0xff >> 4)))
	// fmt.Printf("%#x\n", 0x0f)
}
