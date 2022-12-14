package main

import (
	"cpf/generator"
	"flag"
	"fmt"
)

func main() {
	opts := parseFlags()

	if opts.region < -1 || opts.region > 9 {
		fmt.Println("região inválida")
		flag.Usage()
		return
	}

	for i := 0; i < opts.num; i++ {
		fmt.Println(generator.GenerateCpf(opts.punctuated, opts.region))
	}
}

type ProgramFlags struct {
	num, region int
	punctuated  bool
}

func parseFlags() ProgramFlags {
	num := flag.Int("n", 1, "quantidade de dados a serem gerados")
	region := flag.Int(
		"r", -1,
		`região emissora do cpf
1 – DF, GO, MS, MT e TO
2 – AC, AM, AP, PA, RO e RR
3 – CE, MA e PI
4 – AL, PB, PE, RN
5 – BA e SE
6 – MG
7 – ES e RJ
8 – SP
9 – PR e SC
0 – RS`,
	)
	punctuated := flag.Bool("p", false, "gerar dados com pontuação")
	flag.Parse()

	return ProgramFlags{*num, *region, *punctuated}
}
