package envparser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/seq/linear"
)

var code3to1 = map[string]byte{
	"ALA": 'A', "ARG": 'R', "ASN": 'N', "ASP": 'D', "CYS": 'C', "GLU": 'E',
	"GLN": 'Q', "GLY": 'G', "HIS": 'H', "ILE": 'I', "LEU": 'L', "LYS": 'K',
	"MET": 'M', "PHE": 'F', "PRO": 'P', "SER": 'S', "THR": 'T', "TRP": 'W',
	"TYR": 'Y', "VAL": 'V'}

func envcode(ss string, acc string) (byte, error) {
	var code byte
	if ss == "C" {
		code = 65 //"A"
	} else if ss == "H" {
		code = 72 //"H"
	} else if ss == "S" {
		code = 83 //"S"
	} else {
		return 32, fmt.Errorf("Codigo de estrutura secundaria invalido: %s\n", code)
	}
	//if acc == E code += 0
	if acc == "P1" {
		code += 1
	} else if acc == "P2" {
		code += 2
	} else if acc == "B1" {
		code += 3
	} else if acc == "B2" {
		code += 4
	} else if acc == "B3" {
		code += 5
	}
	return code, nil
}

func GetEnv3D(fn string) (*linear.Seq, *linear.Seq, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao ler o arquivo Environments: %s\n", err)
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return nil, nil, fmt.Errorf("Erro ao ler o arquivo Environments: %s\n", fn)
	}

	header_pos := 0
	for i, l := range lines {

		if len(l) > 3 && l[4:8] == "ResN" {
			header_pos = i
			break
		}
	}
	if header_pos == 0 {
		fmt.Println("problema ao procurar o cabecalho")
		return nil, nil, fmt.Errorf("Erro ao procurar o inicio dos dados no arquivo Environments. Problema no cabeçalho\n")
	}

	col_aa := make([]byte, len(lines)-header_pos-2)
	col_env := make([]byte, len(lines)-header_pos-2)

	for i, l := range lines[header_pos+1 : len(lines)] {
		if len(l) > 30 {
			col_aa[i] = byte(code3to1[l[10:13]])

			env, err := envcode(l[29:30], l[32:34])
			if err != nil {
				fmt.Println("Erro no codigo do environment")
			}
			col_env[i] = env

		}
	}

	id := strings.Split(fn, "/")
	aa_id := fmt.Sprintf("%s |AA|ENV| File %s", id[len(id)-1], fn)
	env_id := fmt.Sprintf("%s |EV|ENV| File %s", id[len(id)-1], fn)

	aa := linear.NewSeq(aa_id, alphabet.BytesToLetters(col_aa), alphabet.Protein)
	env := linear.NewSeq(env_id, alphabet.BytesToLetters(col_env), alphabet.Protein)

	return aa, env, nil
}
