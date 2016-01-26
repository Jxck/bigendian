package bigendian

import (
	"encoding/binary"
	"io"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

var p = log.Println

var o = binary.BigEndian

// delegate
func Read(r io.Reader, data interface{}) error {
	return binary.Read(r, o, data)
}

func Read8(r io.Reader, n *uint8) (err error) {
	err = binary.Read(r, o, n)
	if err != nil {
		return
	}
	return
}

func Read16(r io.Reader, n *uint16) (err error) {
	err = binary.Read(r, o, n)
	if err != nil {
		return
	}
	return
}

func Read24(r io.Reader, n *uint32) (err error) {
	var b [3]byte
	err = binary.Read(r, o, &b)
	if err != nil {
		return
	}

	*n = *n + uint32(b[0])<<16
	*n = *n + uint32(b[1])<<8
	*n = *n + uint32(b[2])

	return
}

func Read32(r io.Reader, n *uint32) (err error) {
	err = binary.Read(r, o, n)
	if err != nil {
		return
	}
	return
}

func Read40(r io.Reader, n *uint64) (err error) {
	var b [5]byte
	err = binary.Read(r, o, &b)
	if err != nil {
		return
	}

	*n = *n + uint64(b[0])<<32
	*n = *n + uint64(b[1])<<24
	*n = *n + uint64(b[2])<<16
	*n = *n + uint64(b[3])<<8
	*n = *n + uint64(b[4])

	return
}

func Read48(r io.Reader, n *uint64) (err error) {
	var b [6]byte
	err = binary.Read(r, o, &b)
	if err != nil {
		return
	}

	*n = *n + uint64(b[0])<<40
	*n = *n + uint64(b[1])<<32
	*n = *n + uint64(b[2])<<24
	*n = *n + uint64(b[3])<<16
	*n = *n + uint64(b[4])<<8
	*n = *n + uint64(b[5])

	return
}

// Read first 1 byte as length
func ReadPrefix1(r io.Reader, b []byte) (err error) {
	var length uint8
	err = binary.Read(r, o, &length)

	var buf []byte = make([]byte, int(length))
	binary.Read(r, o, buf)

  p(b, buf)
	copy(b, buf)
  p(b)
	return
}

// Read first 2 byte as length
func ReadPrefix2(r io.Reader, b []byte) (err error) {
	var length uint16
	err = binary.Read(r, o, &length)

	var buf []byte = make([]byte, int(length))
	binary.Read(r, o, buf)
	copy(b, buf)
	return
}

//////////////////

// delegate
func Write(w io.Writer, data interface{}) error {
	return binary.Write(w, o, data)
}

func Write8(w io.Writer, n uint8) (err error) {
	err = binary.Write(w, o, n)
	if err != nil {
		return
	}
	return
}

func Write16(w io.Writer, n uint16) (err error) {
	err = binary.Write(w, o, n)
	if err != nil {
		return
	}
	return
}

func Write24(w io.Writer, n uint32) (err error) {
	var b [3]byte = [3]byte{
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}

	err = binary.Write(w, o, b)
	if err != nil {
		return
	}

	return
}

func Write32(w io.Writer, n uint32) (err error) {
	err = binary.Write(w, o, n)
	if err != nil {
		return
	}
	return
}

func Write40(w io.Writer, n uint64) (err error) {
	var b [5]byte = [5]byte{
		byte(n >> 32),
		byte(n >> 24),
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}

	err = binary.Write(w, o, b)
	if err != nil {
		return
	}

	return
}

func Write48(w io.Writer, n uint64) (err error) {
	var b [6]byte = [6]byte{
		byte(n >> 40),
		byte(n >> 32),
		byte(n >> 24),
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}

	err = binary.Write(w, o, b)
	if err != nil {
		return
	}

	return
}

// Write first 1 byte as length
func WritePrefix1(w io.Writer, b []byte) (err error) {
  p(len(b), b)
	var length uint8 = uint8(len(b))
	err = binary.Write(w, o, length)
	err = binary.Write(w, o, b)
  p(err)
	return
}

// Write first 2 byte as length
func WritePrefix2(w io.Writer, b []byte) (err error) {
	var length uint16 = uint16(len(b))
	err = binary.Write(w, o, length)
	err = binary.Write(w, o, b)
	return
}
