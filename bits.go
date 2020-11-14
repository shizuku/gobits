package gobits

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Bits :
type Bits struct {
	data   []byte
	offset int
}

// New :
func New() *Bits {
	return &Bits{
		data:   []byte{0},
		offset: 0,
	}
}

// FromBytes :
func FromBytes(bys []byte) *Bits {
	return &Bits{
		data:   bys,
		offset: 0,
	}
}

// FromByte :
func FromByte(by byte, offset int) *Bits {
	if offset < 0 || offset >= 8 {
		log.Fatal("offset out of range.")
	}
	return &Bits{
		data:   []byte{by & (0xFF >> (8 - offset))},
		offset: offset,
	}
}

// Len :
func (b *Bits) Len() int {
	if b.offset == 0 {
		return len(b.data) * 8
	}
	return len(b.data)*8 + b.offset - 8
}

// Offset :
func (b *Bits) Offset() int {
	return b.offset
}

// SetOffset :
func (b *Bits) SetOffset(ofs int) {
	b.offset = ofs
	b.data[len(b.data)-1] = (b.data[len(b.data)-1] & (0xff >> (8 - ofs)))
}

// Append :
func (b *Bits) Append(bts *Bits) {
	l := len(bts.data)
	for i, v := range bts.data {
		if i == l-1 {
			b.AppendBits(v, bts.offset)
		} else {
			b.AppendBits(v, 8)
		}
	}
}

// AppendBits :
func (b *Bits) AppendBits(by byte, offset int) {
	if offset <= 0 || offset > 8 {
		log.Fatal("offset out of range.")
	}
	if b.offset == 0 {
		ny := by & (0xff >> (8 - offset))
		b.data = append(b.data, ny)
		b.offset = b.offset + offset
		if b.offset == 8 {
			b.offset = 0
		}
	} else {
		if b.offset+offset > 8 {
			ey := (by >> (offset - 8 + b.offset)) & (0xff >> b.offset)
			zy := (b.data[len(b.data)-1] << (8 - b.offset)) & 0xff
			ny := by & (0xff >> (8 - offset + 8 - b.offset))
			b.data[len(b.data)-1] = ey | zy
			b.data = append(b.data, ny)
			b.offset = b.offset + offset - 8
		} else {
			ey := by & (0xff >> (8 - offset))
			zy := ((b.data[len(b.data)-1] << (offset)) & 0xff)
			//fmt.Printf("d:%#08b,%#08b\n", ey, zy)
			b.data[len(b.data)-1] = ey | zy
			b.offset = b.offset + offset
			if b.offset == 8 {
				b.offset = 0
			}
		}
	}
}

// AppendBit :
func (b *Bits) AppendBit(by byte) {
	b.AppendBits(by, 1)
}

// AppendByte :
func (b *Bits) AppendByte(by byte) {
	b.AppendBits(by, 8)
}

// Bytes :
func (b *Bits) Bytes() []byte {
	return b.data
}
func (b *Bits) String() string {
	var s strings.Builder
	l := len(b.data)
	for i, v := range b.data {
		if i == l-1 {
			d := 8
			if b.offset != 0 {
				d = b.offset
			}
			s.WriteString(fmt.Sprintf("%0"+fmt.Sprint(d)+"b", v))
		} else {
			s.WriteString(fmt.Sprintf("%08b_", v))
		}
	}
	return s.String()
}

// Itor :
func (b *Bits) Itor() *Iterator {
	return &Iterator{bts: b, idx: 0, offset: 0}
}

// Iterator :
type Iterator struct {
	bts    *Bits
	idx    int
	offset int
}

// Next :
func (it *Iterator) Next() (by byte, idx int, err error) {
	if it.idx > (len(it.bts.data)-1) || (it.idx == (len(it.bts.data)-1) && it.offset >= it.bts.offset) {
		by = 0
		idx = 0
		err = errors.New("end of iterator")
		return
	}
	x := it.bts.data[it.idx]
	if it.idx == (len(it.bts.data) - 1) {
		by = ((x >> (it.bts.offset - it.offset - 1)) & 0x01)
	} else {
		by = ((x >> (7 - it.offset)) & 0x01)
	}
	idx = it.idx*8 + (it.offset)
	err = nil
	it.offset = it.offset + 1
	if it.offset == 8 {
		it.idx = it.idx + 1
		it.offset = 0
	}
	return
}
