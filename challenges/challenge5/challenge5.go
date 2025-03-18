package main

import "fmt"

func somar(n ...int) int {
	total := 0
	for _, numero := range n {
		total += numero
	}
	return total
}

func multiplicar(n ...int) int {
	total := 1
	for _, numero := range n {
		total = total * numero
	}
	return total
}

func subtrair(n ...int) int {
	total := 0
	for _, numero := range n {
		total = numero - total
	}
	return total
}

func dividir(n ...int) int {
	total := 1
	for _, numero := range n {
		total = numero / total
	}
	return total
}

func main() {

	a := somar(1, 3, 5)
	b := multiplicar(4, 5)
	c := subtrair(3, 10)
	d := dividir(2, 10)
	fmt.Println(a, b, c, d)

}
