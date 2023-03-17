package main

import "fmt"

func main() {
	var base float64
	var altura float64
	var profundidade float64
	fmt.Println("Digite a base:")
	fmt.Scan(&base)
	fmt.Println("Digite a altura:")
	fmt.Scan(&altura)
	fmt.Println("Digite a profundidade:")
	fmt.Scan(&profundidade)
	volume := base * altura * profundidade
	fmt.Println("O volume é igual a ", volume)
}
