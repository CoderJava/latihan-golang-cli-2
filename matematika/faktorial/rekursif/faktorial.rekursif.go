package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input angka: ")
	scanner.Scan()
	inputAngka := scanner.Text()
	angka, err := strconv.Atoi(inputAngka)
	if err != nil {
		fmt.Print("Input tidak valid.")
		return
	}

	faktorial := hitungFaktorial(angka)
	fmt.Printf("%d! = %d", angka, faktorial)
}

func hitungFaktorial(angka int) int {
	if angka > 1 {
		return angka * hitungFaktorial(angka-1)
	} else {
		return 1
	}
}
