package main

import "fmt"

func main() {
	divisors := map[int]string{
		3: "Pin",
		5: "Pan",
	}

	for i := 1; i <= 100; i++ {
		output := ""
		for key, value := range divisors {
			if i%key == 0 {
				output += value
			}
		}
		if output == "" {
			output = fmt.Sprint(i)
		}
		fmt.Println(output)
	}
}
