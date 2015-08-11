package ssparser

import (
	//"fmt"
	"testing"
	//"github.com/biogo/biogo/seq/linear"
)

type testdata struct {
	filename string
	aa       string
	ss       string
}

var tests = []testdata{
	{"tests/dssp/2ookA00.dssp", "DKKHGLIGINRIESVFVLKAGTLTHEDYLVITPLEGALSQVDQPKVSLFLDATELDGWDLRAAWDDLKLGLKHKSFRVAILGNKDWQEWAAKIGSWFIAGEIKYFEDEDDALKWLRY", "CCCCCCCCCCCBTTBCCCBCCEECHHHHHHHCCCHHHHTTCCCSSCCEEEECTTCCEECTTCGGGGCCCCCTTCCCCEEEECCSSCCTTTTTGGGGCCCSCEEEESCHHHHHHHHHC"},
}

func TestGetDSSP(t *testing.T) {
	for _, test := range tests {
		dssp_aa, dssp_ss, _ := GetDSSP(test.filename)
		if dssp_aa.Len() != dssp_ss.Len() {
			t.Error("Sequência de aa e sequência de ss com tamanhos diferentes.")
		}
		if dssp_aa.String() != test.aa {
			t.Error("Sequência de aminoácidos diferentes.")
		}
		if dssp_ss.String() != test.ss {
			t.Error("Estrutura secundária diferentes.")
		}
	}
}
