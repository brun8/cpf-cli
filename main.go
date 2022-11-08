package main

import (
	"flag"
	"fmt"
  "cpf/generator"
)

func main() {
	num := flag.Int("n", 1, "quantidade de dados a serem gerados")
	punctuated := flag.Bool("p", false, "gerar dados com pontuação")
	flag.Parse()

	for i := 0; i < *num; i++ {
		fmt.Println(generator.GenerateCpf(*punctuated))
	}
}


