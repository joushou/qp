package qp

import (
	"reflect"
	"testing"
)

// Test if the types live up to their interface
var (
	_ Message = (*SessionRequestDote)(nil)
	_ Message = (*SessionResponseDote)(nil)
	_ Message = (*SimpleReadRequestDote)(nil)
	_ Message = (*SimpleReadResponseDote)(nil)
	_ Message = (*SimpleWriteRequestDote)(nil)
	_ Message = (*SimpleWriteResponseDote)(nil)
)

var MessageTestDataDote = []MessageTestEntry{
	{
		&SessionRequestDote{
			Tag: 45,
			Key: [8]byte{1, 2, 3, 4, 5, 6, 7, 8},
		},
		[]byte{0x2d, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
		[]byte{0xf, 0x0, 0x0, 0x0, 0x96, 0x2d, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
	}, {
		&SessionResponseDote{
			Tag: 45,
		},
		[]byte{0x2d, 0x0},
		[]byte{0x7, 0x0, 0x0, 0x0, 0x97, 0x2d, 0x0},
	}, {
		&SimpleReadRequestDote{
			Tag: 45,
			Fid: 234135,
			Names: []string{
				"Hello",
				"something",
				"wee.exe",
			},
		},
		[]byte{0x2d, 0x0, 0x97, 0x92, 0x3, 0x0, 0x3, 0x0, 0x5, 0x0, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x9, 0x0, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x7, 0x0, 0x77, 0x65, 0x65, 0x2e, 0x65, 0x78, 0x65},
		[]byte{0x28, 0x0, 0x0, 0x0, 0x98, 0x2d, 0x0, 0x97, 0x92, 0x3, 0x0, 0x3, 0x0, 0x5, 0x0, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x9, 0x0, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x7, 0x0, 0x77, 0x65, 0x65, 0x2e, 0x65, 0x78, 0x65},
	}, {
		&SimpleReadResponseDote{
			Tag:  45,
			Data: []byte("something something something"),
		},
		[]byte{0x2d, 0x0, 0x1d, 0x0, 0x0, 0x0, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67},
		[]byte{0x28, 0x0, 0x0, 0x0, 0x99, 0x2d, 0x0, 0x1d, 0x0, 0x0, 0x0, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67},
	}, {
		&SimpleWriteRequestDote{
			Tag: 45,
			Fid: 132412,
			Names: []string{
				"Hello",
				"something",
				"wee.exe",
			},
			Data: []byte("asældfkjasældkgjaæsldkfj"),
		},
		[]byte{0x2d, 0x0, 0x3c, 0x5, 0x2, 0x0, 0x3, 0x0, 0x5, 0x0, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x9, 0x0, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x7, 0x0, 0x77, 0x65, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x1b, 0x0, 0x0, 0x0, 0x61, 0x73, 0xc3, 0xa6, 0x6c, 0x64, 0x66, 0x6b, 0x6a, 0x61, 0x73, 0xc3, 0xa6, 0x6c, 0x64, 0x6b, 0x67, 0x6a, 0x61, 0xc3, 0xa6, 0x73, 0x6c, 0x64, 0x6b, 0x66, 0x6a},
		[]byte{0x47, 0x0, 0x0, 0x0, 0x9a, 0x2d, 0x0, 0x3c, 0x5, 0x2, 0x0, 0x3, 0x0, 0x5, 0x0, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x9, 0x0, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x7, 0x0, 0x77, 0x65, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x1b, 0x0, 0x0, 0x0, 0x61, 0x73, 0xc3, 0xa6, 0x6c, 0x64, 0x66, 0x6b, 0x6a, 0x61, 0x73, 0xc3, 0xa6, 0x6c, 0x64, 0x6b, 0x67, 0x6a, 0x61, 0xc3, 0xa6, 0x73, 0x6c, 0x64, 0x6b, 0x66, 0x6a},
	}, {
		&SimpleWriteResponseDote{
			Tag:   45,
			Count: 12341234,
		},
		[]byte{0x2d, 0x0, 0xf2, 0x4f, 0xbc, 0x0},
		[]byte{0xb, 0x0, 0x0, 0x0, 0x9b, 0x2d, 0x0, 0xf2, 0x4f, 0xbc, 0x0},
	},
}

func TestUnmarshalErrorDote(t *testing.T) {
	for i, tt := range MessageTestDataDote {
		r := reflect.New(reflect.ValueOf(tt.input).Elem().Type()).Interface().(binaryBothWayer)
		testUnmarshal(t, i, r, tt.reference[:len(tt.reference)-1])
	}
}

// This test does NOT guarantee proper 9P2000 spec coding, but ensures at least
// that all codecs are compatible with themselves.
func TestReencodeDote(t *testing.T) {
	for i, tt := range MessageTestDataDote {
		reencode(i, tt.input, tt.reference, t, NineP2000Dote)
	}
}
