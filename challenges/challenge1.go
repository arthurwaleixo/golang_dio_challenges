package main

import "fmt"

const boilingPointInKelvin = 373.15

func main() {
	kelvinTemp := boilingPointInKelvin
	celsiusTemp := kelvinTemp - 273.15

	fmt.Printf("A temperatura de ebulição da água em Kelvin é igual a %g. E a temperatura da ebulição da água em graus Celsius é igual a %g.", kelvinTemp, celsiusTemp)
}
