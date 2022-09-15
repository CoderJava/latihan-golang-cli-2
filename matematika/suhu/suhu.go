package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input suhu celcius: ")
	scanner.Scan()
	inputCelcius := scanner.Text()
	celcius, err := strconv.ParseFloat(inputCelcius, 64)
	if err != nil {
		fmt.Print("Input tidak valid")
		return
	}

	fahrenheit := 1.8*celcius + 32
	kelvin := celcius + 273.15
	reamur := 0.8 * celcius

	formatCelcius := strconv.FormatFloat(celcius, 'f', -1, 64)
	formatFahrenheit := strconv.FormatFloat(fahrenheit, 'f', -1, 64)
	formatKelvin := strconv.FormatFloat(kelvin, 'f', -1, 64)
	formatReamur := strconv.FormatFloat(reamur, 'f', -1, 64)

	fmt.Printf("%s derajat Celcius = %s derajat Fahrenheit\n", formatCelcius, formatFahrenheit)
	fmt.Printf("%s derajat Celcius = %s derajat Kelvin\n", formatCelcius, formatKelvin)
	fmt.Printf("%s derajat Celcius = %s derajat Reamur", formatCelcius, formatReamur)
}
