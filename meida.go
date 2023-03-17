package main

import "fmt"

func main() {
	var x, y, z, a float64
	fmt.Println("Numero 1: ")
	fmt.Scan(&x)
	fmt.Println("Numero 2: ")
	fmt.Scan(&y)
	fmt.Println("Numero 3: ")
	fmt.Scan(&z)
	fmt.Println("Numero 4: ")
	fmt.Scan(&a)
	media := (x + y + z + a) / 4
	fmt.Println("Media: ", media)
}
