package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input panjang persegi: ")
	scanner.Scan()
	strInputPanjangPersegi := scanner.Text()
	inputPanjangPersegi, err := strconv.ParseFloat(strInputPanjangPersegi, 64)
	if err != nil {
		fmt.Println("Input panjang persegi tidak valid.")
		return
	}

	fmt.Print("Input lebar persegi: ")
	scanner.Scan()
	strInputLebarPersegi := scanner.Text()
	inputLebarPersegi, err := strconv.ParseFloat(strInputLebarPersegi, 64)
	if err != nil {
		fmt.Println("Input lebar persegi tidak valid.")
		return
	}

	luasPersegiPanjang := inputPanjangPersegi * inputLebarPersegi
	fmt.Printf("Luas persegi panjang = %.2f", luasPersegiPanjang)
}