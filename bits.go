package gobits

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Bits is a bit array.
//
// data store bits in byte array.
// The last byte may be not full, which empty bits in high(left) palce and NOT empty bits in low(right).
//
// offset is the amount of empty bit in last byte of data.
type Bits struct {
	data   []byte
	offset int
}

// New create an empty Bits.
func New() *Bits {
	return &Bits{
		data:   []byte{},
		offset: 0,
	}
}

// FromBytes create Bits from byte array.
// offset is the amount of empty bit of last byte.
// NOT empty bit of last byte must be the low(right).
func FromBytes(bys []byte, offset int) *Bits {
	return &Bits{
		data:   bys,
		offset: offset,
	}
}

// FromByte create Bits from a byte.
// num is the amount of NOT empty bit.
func FromByte(by byte, num int) *Bits {
	if num < 0 || num > 8 {
		log.Fatal("offset out of range.")
	}
	return &Bits{
		data:   []byte{by & (0xFF >> num)},
		offset: 8 - num,
	}
}

// Len return the length on the basis of bit.
func (b *Bits) Len() int {
	return len(b.data)*8 - b.offset
}

// Offset returns offset which means empty bit of last byte.
func (b *Bits) Offset() int {
	return b.offset
}

// Append append Bits.
func (b *Bits) Append(bts *Bits) {
	l := len(bts.data)
	for i, v := range bts.data {
		if i == l-1 {
			b.AppendBits(v, 8-bts.offset)
		} else {
			b.AppendBits(v, 8)
		}
	}
}

// AppendBits append the low(right) num bit of byte.
func (b *Bits) AppendBits(by byte, num int) {
	if num == 0 {
		return
	}
	if num < 0 || num > 8 {
		log.Fatal("offset out of range.")
	}
	if b.offset == 0 {
		b.data = append(b.data, 0)
		b.offset = 8
	}
	if b.offset >= num {
		rby := by & (0xff >> (8 - num))
		zby := (b.data[len(b.data)-1] << num) & 0xff
		b.data[len(b.data)-1] = rby | zby
		b.offset = b.offset - num
	} else {
		rby := (by >> (num - b.offset)) & (0xff >> (8 - b.offset))
		zby := (b.data[len(b.data)-1] << b.offset) & 0xff
		nby := by & (0xff >> (8 - num + b.offset))
		b.data[len(b.data)-1] = rby | zby
		b.data = append(b.data, nby)
		b.offset = 8 - num + b.offset
	}

	// if b.offset == 0 {
	// 	ny := by & (0xff >> (8 - num))
	// 	b.data = append(b.data, ny)
	// 	b.offset = b.offset + num
	// 	if b.offset == 8 {
	// 		b.offset = 0
	// 	}
	// } else {
	// 	if 8-b.offset+num > 8 {
	// 		ey := (by >> (num - 8 + b.offset)) & (0xff >> b.offset)
	// 		zy := (b.data[len(b.data)-1] << (8 - b.offset)) & 0xff
	// 		ny := by & (0xff >> (8 - num + 8 - b.offset))
	// 		b.data[len(b.data)-1] = ey | zy
	// 		b.data = append(b.data, ny)
	// 		b.offset = b.offset + num - 8
	// 	} else {
	// 		ey := by & (0xff >> (8 - num))
	// 		zy := ((b.data[len(b.data)-1] << (num)) & 0xff)
	// 		//fmt.Printf("d:%#08b,%#08b\n", ey, zy)
	// 		b.data[len(b.data)-1] = ey | zy
	// 		b.offset = b.offset + num
	// 		if b.offset == 8 {
	// 			b.offset = 0
	// 		}
	// 	}
	// }
}

// AppendBit append the low(right) 1 bit of a byte.
func (b *Bits) AppendBit(by byte) {
	b.AppendBits(by, 1)
}

// AppendByte append a totle byte.
func (b *Bits) AppendByte(by byte) {
	b.AppendBits(by, 8)
}

// Bytes returns the data in byte array.
// The empty bit of last byte will be set to 0.
func (b *Bits) Bytes() []byte {
	return b.data
}

// String returns string of 0 or 1.
// There will be a '_' between each byte.
func (b *Bits) String() string {
	var s strings.Builder
	it := b.Itor()
	for {
		ch, idx, err := it.Next()
		if err != nil {
			break
		}
		if idx%8 == 7 {
			s.WriteString(fmt.Sprintf("%d_", ch))
		} else {
			s.WriteString(fmt.Sprintf("%d", ch))
		}
	}
	return s.String()
}

// Itor create an iterator for current Bits.
func (b *Bits) Itor() *Iterator {
	return &Iterator{bts: b, idx: 0, pos: 0}
}

// Iterator is the iterator for Bits which iterate Bits by bit.
//
// pos count 0 to 8 form high(left) to low(right)
type Iterator struct {
	bts *Bits
	idx int
	pos int
}

// Next returns next bit fom iterator.
// err will be not nil when end.
func (it *Iterator) Next() (by byte, idx int, err error) {
	if it.idx > (len(it.bts.data)-1) || (it.idx == (len(it.bts.data)-1) && it.pos >= (8-it.bts.offset)) {
		by = 0
		idx = 0
		err = errors.New("end of iterator")
		return
	}
	x := it.bts.data[it.idx]
	if it.idx == (len(it.bts.data) - 1) {
		by = ((x >> (8 - it.bts.offset - it.pos - 1)) & 1)
	} else {
		by = ((x >> (7 - it.pos)) & 1)
	}
	idx = it.idx*8 + (it.pos)
	err = nil
	it.pos = it.pos + 1
	if it.pos == 8 {
		it.idx = it.idx + 1
		it.pos = 0
	}
	return
}
