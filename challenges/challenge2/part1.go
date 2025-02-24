package main

import "fmt"

func main() {
	value := 3

	for i := 1; i <= 100; i++ {
		if i%value == 0 {
			fmt.Println(i)
		}
	}
}
