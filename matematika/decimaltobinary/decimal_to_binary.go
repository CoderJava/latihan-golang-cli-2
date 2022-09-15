package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input angka desimal: ")
	scanner.Scan()
	inputAngkaDesimal := scanner.Text()
	angkaDesimal, err := strconv.Atoi(inputAngkaDesimal)
	if err != nil {
		fmt.Print("Input tidak valid")
		return
	}

	biner := konversiDesimalKeBiner(angkaDesimal)
	fmt.Printf("Angka binernya adalah: %s", biner)
}

func konversiDesimalKeBiner(angkaDesimal int) (biner string) {
	var binerTemp []int
	for {
		sisa := angkaDesimal % 2
		binerTemp = append(binerTemp, sisa)
		angkaDesimal = angkaDesimal / 2
		if angkaDesimal == 0 {
			break
		}
	}

	for index := len(binerTemp) - 1; index >= 0; index-- {
		biner += strconv.Itoa(binerTemp[index])
	}
	return
}
