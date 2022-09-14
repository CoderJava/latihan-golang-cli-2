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
		fmt.Print("Input tidak valid")
		return
	}

	hasil := 1
	for index := 2; index <= angka; index++ {
		hasil *= index
	}

	fmt.Printf("%d! = %d", angka, hasil)
}
