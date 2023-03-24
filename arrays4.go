package main

import "fmt"

func main() {
	var n1, n2, n3, n4, n5, n6 float64
	var array []float64
	var sum float64
	fmt.Println("primeiro numero: ")
	fmt.Scan(&n1)
	fmt.Println("primeiro segundo: ")
	fmt.Scan(&n2)
	fmt.Println("primeiro terceiro: ")
	fmt.Scan(&n3)
	fmt.Println("primeiro quarto: ")
	fmt.Scan(&n4)
	fmt.Println("primeiro quinto: ")
	fmt.Scan(&n5)
	fmt.Println("primeiro sexto: ")
	fmt.Scan(&n6)
	array = append(array, n1, n2, n3, n4, n5, n6)
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	media := sum / 6
	fmt.Println(media)
}
