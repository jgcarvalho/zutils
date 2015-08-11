package ssparser

import (
	"github.com/biogo/biogo/alphabet"
	//"github.com/biogo/biogo/seq"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/biogo/biogo/seq/linear"
)

func GetStride(fn string) (*linear.Seq, *linear.Seq, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao ler o arquivo Stride: %s\n", err)
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return nil, nil, fmt.Errorf("Erro ao ler o arquivo Stride: %s\n", fn)
	}

	seq_aa := ""
	seq_ss := ""
	for _, l := range lines {
		if len(l) > 3 && l[0:3] == "SEQ" {
			seq_aa += l[10:60]
		} else if len(l) > 3 && l[0:3] == "STR" {
			seq_ss += l[10:60]
		}
	}

	if seq_aa == "" || seq_ss == "" {
		return nil, nil, fmt.Errorf("Sequencia de aa ou estrutura secundaria não encontrados no arquivo Stride: %s\n", fn)
	}

	seq_aa = strings.TrimSpace(seq_aa)
	seq_ss = seq_ss[:len(seq_aa)]
	seq_ss = strings.Replace(seq_ss, " ", "C", -1)

	id := strings.Split(fn, "/")
	aa_id := fmt.Sprintf("%s |AA|STRIDE| File %s", id[len(id)-1], fn)
	ss_id := fmt.Sprintf("%s |SS|STRIDE| File %s", id[len(id)-1], fn)

	aa := linear.NewSeq(aa_id, alphabet.BytesToLetters([]byte(seq_aa)), alphabet.Protein)
	ss := linear.NewSeq(ss_id, alphabet.BytesToLetters([]byte(seq_ss)), alphabet.Protein)

	return aa, ss, nil
}
