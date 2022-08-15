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
		fmt.Println("Input jumlah element tidak valid.")
		return
	}

	fmt.Printf("Input %d angka (dipisah dengan enter):", inputJumlahElement)
	fmt.Println()
	elements := make([]int, inputJumlahElement)
	var average float64
	var sum int
	for x := 0; x < inputJumlahElement; x++ {
		scanner.Scan()
		strInputElement := scanner.Text()
		inputElement, err := strconv.Atoi(strInputElement)
		if err != nil {
			fmt.Println("Input angka tidak valid.")
			return
		}
		elements[x] = inputElement
		sum += inputElement
	}
	average = float64(sum) / float64(len(elements))
	fmt.Printf("Nilai rata-rata dari %d input adalah: %.2f", inputJumlahElement, average)
}
