package main

import "fmt"

func main() {
	fmt.Println(convertTemperature(36.50))
}

func convertTemperature(celsius float64) []float64 {
    kelvin := celsius + 273.15
    fahrenheit := celsius * 1.80 + 32.00

    result := []float64 {}
    result = append(result, kelvin, fahrenheit)
    return result
}