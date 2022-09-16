package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Format persamaan: ax^2 + bx + c = 0")

	fmt.Print("Input nilai a: ")
	scanner.Scan()
	inputNilaiA := scanner.Text()
	nilaiA, err := strconv.Atoi(inputNilaiA)
	if err != nil {
		fmt.Print("Input tidak valid")
		return
	}

	fmt.Print("Input nilai b: ")
	scanner.Scan()
	inputNilaiB := scanner.Text()
	nilaiB, err := strconv.Atoi(inputNilaiB)
	if err != nil {
		fmt.Print("Input tidak valid")
		return
	}

	fmt.Print("Input nilai c: ")
	scanner.Scan()
	inputNilaiC := scanner.Text()
	nilaiC, err := strconv.Atoi(inputNilaiC)
	if err != nil {
		fmt.Print("Input tidak valid")
		return
	}
	fmt.Println()

	d := hitungDeterminan(nilaiA, nilaiB, nilaiC)
	fmt.Printf("Diskriminan = %d", d)

	switch {
	case d > 0:
		akar := math.Sqrt(float64(d))
		x1 := (float64(-nilaiB) + akar) / float64(2*nilaiA)
		x2 := (float64(-nilaiB) - akar) / float64(2*nilaiA)

		// 36 / 10

		formatX1 := strconv.FormatFloat(x1, 'f', -1, 64)
		formatX2 := strconv.FormatFloat(x2, 'f', -1, 64)
		fmt.Println(" (akar real dan berbeda)")
		fmt.Printf("x1 = %s\n", formatX1)
		fmt.Printf("x2 = %s", formatX2)
	case d < 0:
		fmt.Print(" (akar tidak real / imajiner)")
	case d == 0:
		akar := math.Sqrt(float64(d))
		x := (float64(-nilaiB) + akar) / float64(2*nilaiA)
		formatX := strconv.FormatFloat(x, 'f', -1, 64)
		fmt.Println(" (akar real dan sama)")
		fmt.Printf("x1 = %s\n", formatX)
		fmt.Printf("x2 = %s\n", formatX)
	}
}

func hitungDeterminan(nilaiA, nilaiB, nilaiC int) (result int) {
	result = (nilaiB * nilaiB) - (4 * nilaiA * nilaiC)
	return
}
