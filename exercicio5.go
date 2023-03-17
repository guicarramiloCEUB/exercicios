package main

import "fmt"

func main() {
	var cambio, dolar float64
	fmt.Println("cotacao dolar: ")
	fmt.Scan(&cambio)
	fmt.Println("valor em dolar: ")
	fmt.Scan(&dolar)
	reais := cambio * dolar
	
}
