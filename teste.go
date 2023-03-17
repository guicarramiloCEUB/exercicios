package main

import "fmt"

func main() {
	var nome string
	var idade int8
	var peso float64
	fmt.Println("digite nome:")
	fmt.Scan(&nome)
	fmt.Println("digite idade:")
	fmt.Scan(&idade)
	fmt.Println("digite peso:")
	fmt.Scan(&peso)
	fmt.Println("Nome:", nome)
	fmt.Println("Idade: ", idade)
	fmt.Print("Peso: ", peso)
}
