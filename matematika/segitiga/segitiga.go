package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input alas segitiga: ")
	scanner.Scan()
	strInputAlas := scanner.Text()
	inputAlas, err := strconv.ParseFloat(strInputAlas, 64)
	if err != nil {
		fmt.Println("Input alas segitiga tidak valid.")
		return
	}

	fmt.Print("Input tinggi segitiga: ")
	scanner.Scan()
	strInputTinggi := scanner.Text()
	inputTinggi, err := strconv.ParseFloat(strInputTinggi, 64)
	if err != nil {
		fmt.Println("Input tinggi segitiga tidak valid.")
		return
	}

	luas := 0.5 * inputAlas * inputTinggi
	fmt.Printf("Luas segitiga: %.2f", luas)
}