package main

import (
	"github.com/biogo/biogo/align"
	"github.com/biogo/biogo/align/matrix"
	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/seq"
	//	"github.com/biogo/biogo/alphabet"
	"flag"
	"fmt"
	//"github.com/biogo/biogo/io/seqio/fasta"
	//	"github.com/biogo/biogo/seq/linear"
	//	"reflect"

	//	"pepplanes"
	"github.com/jgcarvalho/zutils/ssparser"
)

func alinha(seq_aa, ss_aa, ss_ss seq.Sequence) (string, string, string) {
	nw := align.NWAffine{
		Matrix:  matrix.MATCH,
		GapOpen: 0,
	}

	//ALINHAMENTO PODE SER DESNECESSARIO E COMO É CUSTOSO, CONVÉM TESTAR ANTES

	//Alinhamento DSSP
	aln, err := nw.Align(seq_aa, ss_aa)
	var f_aa, f_ss [2]alphabet.Slice
	if err == nil {
		// fmt.Printf("%s\n", aln)
		f_aa = align.Format(seq_aa, ss_aa, aln, '_')
		f_ss = align.Format(seq_aa, ss_ss, aln, '_')
		// fmt.Printf("%s\n%s\n%s\n", f_aa[0], f_aa[1], f_ss[1])
	} else {
		fmt.Println("O ERRO E:", err)
	}
	return fmt.Sprint(f_aa[0]), fmt.Sprint(f_aa[1]), fmt.Sprint(f_ss[1])
}

func main() {

	fastaFName := flag.String("fa", "", "fasta file")
	// // pdbFName := flag.String("pdb", "", "pdb file")
	dsspFName := flag.String("dssp", "", "dssp file")
	strideFName := flag.String("stride", "", "stride file")
	// // doKaksi := flag.Bool("kaksi", false, "kaksi")
	// // doBba := flag.Bool("bba", false, "bba")
	// // classBba := flag.Int("class_bba", 0, "class bba")
	// // classAa := flag.Int("class_aa", 0, "class aa")

	flag.Parse()

	var seq_aa,
		dssp_aa,
		dssp_ss,
		stride_aa,
		stride_ss seq.Sequence
	var err error

	// seq_aa, err := ssparser.GetFasta("/home/jgcarvalho/sync/data/seqdb/cath-s35/2ookA00.fasta")
	// // seq_aa, err := ssparser.GetFasta(*fastaFName)
	// if err != nil {
	// 	fmt.Println("Erro no processamento do Fasta", err)
	// }
	dssp_aa, dssp_ss, err = ssparser.GetDSSP("/home/jgcarvalho/sync/data/ssdb-tmp/cath-s35/dssp/2ookA00.dssp")
	// dssp_aa, dssp_ss, err = ssparser.GetDSSP(*dsspFName)
	if err != nil {
		fmt.Println("Erro no processamento do DSSP", err)
	}

	// Le a sequencia fasta
	if *fastaFName == "" {
		fmt.Println("É necessario o arquivo fasta")
		return
	} else {
		seq_aa, err = ssparser.GetFasta("/home/jgcarvalho/sync/data/seqdb/cath-s35/2ookA00.fasta")
		// seq_aa, err := ssparser.GetFasta(*fastaFName)
		if err != nil {
			fmt.Println("Erro no processamento do Fasta", err)
		}
	}

	// OPCOES
	// Le o arquivo dssp (aa e ss)
	if *dsspFName == "" {
		fmt.Println("É necessario o arquivo dssp")
		return
	} else {
		dssp_aa, dssp_ss, err = ssparser.GetDSSP("/home/jgcarvalho/sync/data/ssdb-tmp/cath-s35/dssp/2ookA00.dssp")
		// dssp_aa, dssp_ss, err = ssparser.GetDSSP(*dsspFName)
		if err != nil {
			fmt.Println("Erro no processamento do DSSP", err)
		}
	}

	// Le o arquivo stride (aa e ss)
	if *strideFName == "" {
		fmt.Println("É necessario o arquivo stride")
		return
	} else {
		stride_aa, stride_ss, err = ssparser.GetStride("/home/jgcarvalho/sync/data/ssdb-tmp/cath-s35/stride/2ookA00.stride")
		// stride_aa, stride_ss, err = ssparser.GetStride(*strideFName)
		if err != nil {
			fmt.Println("Erro no processamento do Stride", err)
		}
	}

	// Le o arquivo kaksi (aa e ss)

	// Le o ambiente (environment Verify 3D) do resíduos
	// env_aa, env, _ := envparser.GetEnv3D("/home/jgcarvalho/sync/data/envdb-tmp/cath-s35/2ookA00.env")

	// Le o arquivo de planos das ligacoes peptidicas
	//pepplanes.GetPlanes("/home/jgcarvalho/sync/data/planesdb-tmp/cath-s35/2ookA00.dat")

	//ALINHAMENTO PODE SER DESNECESSARIO E COMO É CUSTOSO, CONVÉM TESTAR ANTES
	aa, dsspaa, dsspss := alinha(seq_aa, dssp_aa, dssp_ss)
	_, strideaa, stridess := alinha(seq_aa, stride_aa, stride_ss)
	fmt.Printf("%s\n%s\n%s\n%s\n%s\n", aa, dsspaa, dsspss, strideaa, stridess)

	// Escreve um arquivo com os dados de ss e gaps em um arquivo JSON
}
