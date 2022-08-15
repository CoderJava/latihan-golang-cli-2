package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input jumlah element array: ")
	scanner.Scan()
	strInputJumlahElement := scanner.Text()
	inputJumlahElement, err := strconv.Atoi(strInputJumlahElement)
	if err != nil {
		fmt.Println("Input jumlah element array tidak valid.")
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
			fmt.Println("Input angka tidak valid.")
			return
		}
		elements[x] = inputElement
	}

	// Cara singkat pakai library
	/* sort.Ints(elements)
	angkaTerbesar := elements[len(elements)-1]
	fmt.Printf("Angka terbesar adalah: %d", angkaTerbesar) */

	// Cara manual
	angkaTerbesar := elements[0]
	for _, element := range elements {
		if element > angkaTerbesar {
			angkaTerbesar = element
		}
	}
	fmt.Printf("Angka terbesar adalah: %d", angkaTerbesar)
}
