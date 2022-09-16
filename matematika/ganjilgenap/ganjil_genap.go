package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input sebuah bilangan bulat: ")
	scanner.Scan()
	inputStrBilanganBulat := scanner.Text()
	inputBilanganBulat, err := strconv.Atoi(inputStrBilanganBulat)
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

	if inputBilanganBulat%2 == 0 {
		fmt.Printf("%d adalah bilangan genap", inputBilanganBulat)
	} else {
		fmt.Printf("%d adalah bilangan ganjil", inputBilanganBulat)
	}
}
