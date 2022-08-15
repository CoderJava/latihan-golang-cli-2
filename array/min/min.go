package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input jumlah element Array: ")
	scanner.Scan()
	strInputJumlahElement := scanner.Text()
	inputJumlahElement, err := strconv.Atoi(strInputJumlahElement)
	if err != nil {
		fmt.Println("Input jumlah element tidak valid.")
		return
	}

	fmt.Printf("Input %d angka (dipisah dengan enter):", inputJumlahElement)
	fmt.Println()
	elements := make([]int, inputJumlahElement)
	for x := 0; x < inputJumlahElement; x++ {
		scanner.Scan()
		strInputElement := scanner.Text()
		inputElement, err := strconv.Atoi(strInputElement)
		if err != nil {
			fmt.Print("Input angka tidak valid.")
			return
		}
		elements[x] = inputElement
	}

	// Cara sederhana pakai library
	/* sort.Ints(elements)
	angkaTerkecil := elements[0]
	fmt.Printf("Angka terkecil adalah: %d", angkaTerkecil) */

	// Cara manual
	angkaTerkecil := elements[0]
	for _, element := range elements {
		if element < angkaTerkecil {
			angkaTerkecil = element
		}
	}
	fmt.Printf("Angka terkecil adalah: %d", angkaTerkecil)
}
