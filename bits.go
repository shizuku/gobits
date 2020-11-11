package gobits

import (
	"fmt"
)

type Bits struct {
	data []byte
	len  int
}

func New() *Bits {
	return &Bits{}
}
func FromBytes(bys []byte) *Bits {
	return &Bits{
		data: bys,
		len:  len(bys) * 8,
	}
}
func FromByte(by byte, nn int) *Bits {
	return &Bits{
		data: []byte{by & (0xFF >> (8 - nn))},
		len:  nn,
	}
}
func (b *Bits) Len() int {
	return b.len
}
func (b *Bits) AppendBits(bts *Bits) error {
	return nil
}
func (b *Bits) AppendByte(by byte) error {
	if b.len%8 == 0 {
		b.data = append(b.data, by)
		b.len = b.len + 8
	} else {

	}
	return nil
}
func (b *Bits) Bytes() []byte {
	return b.data
}
func (b *Bits) Print() {
	fmt.Print("0b")
	for _, v := range b.data {
		fmt.Printf("%b_", v)
	}
}
