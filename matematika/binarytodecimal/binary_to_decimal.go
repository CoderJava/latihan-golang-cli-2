package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input angka biner: ")
	scanner.Scan()
	inputAngkaBiner := scanner.Text()

	desimal := konversiBinerKeDesimal(inputAngkaBiner)
	fmt.Printf("Angka desimal dari biner %s adalah = %d", inputAngkaBiner, desimal)
}

func konversiBinerKeDesimal(biner string) int {
	var desimal int
	listBiners := strings.Split(biner, "")
	lenBiner := len(listBiners) - 1
	for index, elementBiner := range listBiners {
		if elementBiner != "0" && elementBiner != "1" {
			return -1
		}
		resultPow := math.Pow(float64(2), float64(lenBiner-index))
		if elementBiner == "1" {
			desimal += int(resultPow)
		}
	}
	return desimal
}
