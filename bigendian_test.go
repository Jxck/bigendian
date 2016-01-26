package bigendian

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func Test(t *testing.T) {
	b := []byte{0x01, 0x02, 0x03, 0x04}
	t.Logf("%#v", b)
	buffer := bytes.NewBuffer(b)
	t.Logf("%#v", buffer.Next(2))

	var n uint8
	err := binary.Read(buffer, binary.BigEndian, &n)
	t.Log(n)

	var m uint16
	err = binary.Read(buffer, binary.BigEndian, &m)
	t.Log(m)

	if err == nil {
		t.Errorf("got %v", err)
	}
}

//var testcase = []struct {
//	in, out string
//}{
//	{"a", "b"},
//	{"a", "b"},
//}
//
//func TestParameter(t *testing.T) {
//	for _, tc := range testcase {
//		in := tc.in
//		out := tc.out
//		if actual != expected {
//			t.Errorf("\ngot  %v\nwant %v", actual, expected)
//		}
//	}
//}
//
