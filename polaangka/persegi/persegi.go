package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input besar perseg: ")
	scanner.Scan()
	strInputBesarPersegi := scanner.Text()
	inputBesarPersegi, err := strconv.Atoi(strInputBesarPersegi)
	if err != nil {
		fmt.Println("Input besar persegi tidak valid.")
		return
	}

	// Varian 1
	for x := 1; x <= inputBesarPersegi; x++ {
		for y := 1; y <= inputBesarPersegi; y++ {
			fmt.Printf(" %d", x)
		}
		fmt.Println()
	}
	fmt.Println()

	// Varian 2
	for x := 1; x <= inputBesarPersegi; x++ {
		for y := 1; y <= inputBesarPersegi; y++ {
			fmt.Printf(" %d", y)
		}
		fmt.Println()
	}
	fmt.Println()

	// Varian 3
	maxValue := inputBesarPersegi * inputBesarPersegi
	counter := 1
	for x := 1; x <= inputBesarPersegi; x++ {
		for y := 1; y <= inputBesarPersegi; y++ {
			if maxValue > 10 {
				if counter < 10 {
					fmt.Printf("  %d ", counter)
				} else {
					fmt.Printf(" %d ", counter)
				}
			} else {
				fmt.Printf(" %d", counter)
			}
			counter++
		}
		fmt.Println()
	}
}
