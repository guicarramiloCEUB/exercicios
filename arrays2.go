package main

import "fmt"

func main() {
	var nomes []string
	var nome1, nome2 string
	fmt.Println("primeiro nome: ")
	fmt.Scan(&nome1)
	fmt.Println("segundo nome: ")
	fmt.Scan(&nome2)
	nomes = append(nomes, nome1, nome2)
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}
}
