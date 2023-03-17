package main

import "fmt"

func main() {
	var base float64
	var altura float64
	fmt.Println("insira a base: ")
	fmt.Scan(&base)
	fmt.Println("insira a altura: ")
	fmt.Scan(&altura)
	area := (base * altura) / 2
	fmt.Println((area))

}
