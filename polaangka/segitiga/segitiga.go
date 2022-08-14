package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input tinggi segitiga: ")
	scanner.Scan()
	strInputTinggiSegitiga := scanner.Text()
	inputTinggiSegitiga, err := strconv.Atoi(strInputTinggiSegitiga)
	if err != nil {
		fmt.Println("Input tinggi segitiga tidak valid.")
		return
	}

	// Varian 1
	for x := 1; x <= inputTinggiSegitiga; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf(" %d", x)
		}
		fmt.Println()
	}
	fmt.Println()

	// Varian 2
	for x := 1; x <= inputTinggiSegitiga; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf(" %d", y)
		}
		fmt.Println()
	}
	fmt.Println()

	// Varian 3
	counter := 1
	for x := 1; x <= inputTinggiSegitiga; x++ {
		for y := 1; y <= x; y++ {
			if counter < 10 {
				fmt.Printf("  %d ", counter)
			} else {
				fmt.Printf(" %d ", counter)
			}
			counter++
		}
		fmt.Println()
	}
}
