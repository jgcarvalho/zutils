package ssparser

import (
	//"fmt"
	"testing"
	//"code.google.com/p/biogo/seq/linear"
)

// type testdata struct {
// 	filename string
// 	aa       string
// 	ss       string
// }

var stridetests = []testdata{
	{"tests/stride/2ookA00.stride", "DKKHGLIGINRIESVFVLKAGTLTHEDYLVITPLEGALSQVDQPKVSLFLDATELDGWDLRAAWDDLKLGLKHKSFRVAILGNKDWQEWAAKIGSWFIAGEIKYFEDEDDALKWLRY", "CCCCCCCCCCTBTTBCCCEEEEECHHHHHHHHCHHHHHHCCCTTTTCEEEEEEEEEEETTTTGGGGCCCCTTTTTCCEEEECCCCCTTTTTTGGGGCCTTTEEEECCHHHHHHHHHC"},
}

func TestGetStride(t *testing.T) {
	for _, test := range stridetests {
		stride_aa, stride_ss, err := GetStride(test.filename)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if stride_aa.Len() != stride_ss.Len() {
			t.Error("Sequência de aa e sequência de ss com tamanhos diferentes.")
		}
		if stride_aa.String() != test.aa {
			t.Error("Sequência de aminoácidos diferentes.")
		}
		if stride_ss.String() != test.ss {
			t.Error("Estrutura secundária diferentes.")
		}
	}
}
