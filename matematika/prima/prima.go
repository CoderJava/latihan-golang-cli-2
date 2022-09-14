package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input satu angka bulat: ")
	scanner.Scan()
	inputAngka := scanner.Text()
	angka, err := strconv.Atoi(inputAngka)
	if err != nil {
		fmt.Print("Input tidak valid")
		return
	}

	if angka == 1 || angka == 2 || angka == 3 {
		fmt.Printf("%d adalah angka prima", angka)
	} else if angka%2 == 0 || angka%3 == 0 {
		fmt.Printf("%d bukan angka prima", angka)
	} else {
		fmt.Printf("%d adalah angka prima", angka)
	}
}
