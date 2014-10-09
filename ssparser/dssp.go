package ssparser

import (
	"code.google.com/p/biogo/alphabet"
	//"code.google.com/p/biogo/seq"
	"code.google.com/p/biogo/seq/linear"
	"fmt"
	"io/ioutil"
	"strings"
)

// GetDSSP recebe a nome do arquivo gerado pelo programa DSSP e extrai os dados
// da sequência de resíduos e da estrutura secundária atribuída.
//
// Os códigos da estrutura secundária do DSSP são:
//
// 	H = alpha helix
// 	B = residue in isolated beta-bridge
// 	E = extended strand, participates in beta ladder
// 	G = 3-helix (3/10 helix)
// 	I = 5 helix (pi helix)
// 	T = hydrogen bonded turn
// 	S = bend
//	C = Loops or irregular (no dssp é um espaço em branco ' ', mas isso é alterado neste programa)
//
// BUG: Não diferencia entre cadeias diferentes do PDB
func GetDSSP(fn string) (aa *linear.Seq, ss *linear.Seq, err error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao ler o arquivo DSSP: %s\n", err)
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return nil, nil, fmt.Errorf("Erro ao ler o arquivo DSSP: %s\n", fn)
	}

	header_pos := 0
	for i, l := range lines {

		if len(l) > 3 && l[2:3] == "#" {
			header_pos = i
			break
		}
	}
	if header_pos == 0 {
		return nil, nil, fmt.Errorf("Erro ao procurar o inicio dos dados no arquivo DSSP. Problema no cabeçalho\n")
	}

	//WHY -2 and not -1 ??!!!!!!
	tmp_aa := make([]byte, len(lines)-header_pos)
	tmp_ss := make([]byte, len(lines)-header_pos)

	i := 0
	for _, l := range lines[header_pos+1 : len(lines)] {
		if len(l) > 120 && l[13] != '!' {
			tmp_aa[i] = l[13]

			//byte 32 == " "
			if l[16] == ' ' {
				//byte 67 == "C"
				tmp_ss[i] = 'C'
			} else {
				tmp_ss[i] = l[16]
			}
			i++
		}
	}

	col_aa := string(tmp_aa[:i])
	col_ss := string(tmp_ss[:i])

	id := strings.Split(fn, "/")
	aa_id := fmt.Sprintf("%s |AA|DSSP| File %s", id[len(id)-1], fn)
	ss_id := fmt.Sprintf("%s |SS|DSSP| File %s", id[len(id)-1], fn)

	aa = linear.NewSeq(aa_id, alphabet.BytesToLetters([]byte(col_aa)), alphabet.Protein)
	ss = linear.NewSeq(ss_id, alphabet.BytesToLetters([]byte(col_ss)), alphabet.Protein)

	return aa, ss, nil
}

//DKKHGLIGINRIESVFVLKAGTLTHEDYLVITPLEGALSQVDQPKVSLFLDATELDGWDLRAAWDDLKLGLKHKSFRVAILGNKDWQEWAAKIGSWFIAGEIKYFEDED
//DKKHGLIGINRIESVFVLKAGTLTHEDYLVITPLEGALSQVDQPKVSLFLDATELDGWDLRAAWDDLKLGLKHKSFRVAILGNKDWQEWAAKIGSWFIAGEIKYFEDEDDALKWLRY
