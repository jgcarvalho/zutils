package main

import (
	"code.google.com/p/biogo/align"
	"code.google.com/p/biogo/align/matrix"
	//	"code.google.com/p/biogo/alphabet"
	"fmt"
	//"code.google.com/p/biogo/io/seqio/fasta"
	//	"code.google.com/p/biogo/seq/linear"
	//	"reflect"
	"bitbucket.org/zgcarvalho/zutils/envparser"
	//	"pepplanes"
	"bitbucket.org/zgcarvalho/zutils/ssparser"
)

func main() {

	// Le a sequencia fasta
	seq_aa, err := ssparser.GetFasta("/home/jgcarvalho/sync/data/seqdb/cath-s35/2ookA00.fasta")
	if err != nil {
		fmt.Println("Erro no processamento do Fasta", err)
	}

	// OPCOES
	// Le o arquivo dssp (aa e ss)
	dssp_aa, dssp_ss, err := ssparser.GetDSSP("/home/jgcarvalho/sync/data/ssdb-tmp/cath-s35/dssp/2ookA00.dssp")
	if err != nil {
		fmt.Println("Erro no processamento do DSSP", err)
	}
	// Le o arquivo stride (aa e ss)
	//stride_aa, stride_ss, err := ssparser.GetStride("/home/jgcarvalho/sync/data/ssdb-tmp/cath-s35/stride/1a1xA00.stride")
	stride_aa, stride_ss, err := ssparser.GetStride("/home/jgcarvalho/sync/data/ssdb-tmp/cath-s35/stride/2ookA00.stride")
	if err != nil {
		fmt.Println("Erro no processamento do Stride", err)
	}
	// Le o arquivo kaksi (aa e ss)

	// Le o ambiente (environment Verify 3D) do resíduos
	env_aa, env, _ := envparser.GetEnv3D("/home/jgcarvalho/sync/data/envdb-tmp/cath-s35/2ookA00.env")

	// Le o arquivo de planos das ligacoes peptidicas
	//pepplanes.GetPlanes("/home/jgcarvalho/sync/data/planesdb-tmp/cath-s35/2ookA00.dat")

	// Alinha o aa com o fasta e transfere os gaps para o ss
	nw := align.NWAffine{
		Matrix:  matrix.MATCH,
		GapOpen: 0,
	}

	//ALINHAMENTO PODE SER DESNECESSARIO E COMO É CUSTOSO, CONVÉM TESTAR ANTES

	//Alinhamento DSSP
	aln, err := nw.Align(seq_aa, dssp_aa)
	if err == nil {
		fmt.Printf("%s\n", aln)
		f_aa := align.Format(seq_aa, dssp_aa, aln, '_')
		f_ss := align.Format(seq_aa, dssp_ss, aln, '_')
		fmt.Printf("%s\n%s\n%s\n", f_aa[0], f_aa[1], f_ss[1])
	} else {
		fmt.Println("O ERRO E:", err)
	}

	//Alinhamento STRIDE
	aln, err = nw.Align(seq_aa, stride_aa)
	if err == nil {
		fmt.Printf("%s\n", aln)
		f_aa := align.Format(seq_aa, stride_aa, aln, '_')
		f_ss := align.Format(seq_aa, stride_ss, aln, '_')
		fmt.Printf("%s\n%s\n%s\n", f_aa[0], f_aa[1], f_ss[1])
	}

	//Alinhamento ENVIRONMENTS
	aln, err = nw.Align(seq_aa, env_aa)
	if err == nil {
		fmt.Printf("%s\n", aln)
		f_aa := align.Format(seq_aa, env_aa, aln, '_')
		f_env := align.Format(seq_aa, env, aln, '_')
		fmt.Printf("%s\n%s\n%s\n", f_aa[0], f_aa[1], f_env[1])
	}
	// Escreve um arquivo com os dados de ss e gaps em um arquivo JSON
}
