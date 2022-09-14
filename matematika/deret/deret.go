package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan tipe jumlah deret [A/B/C/D]: ")
	scanner.Scan()
	inputTipe := strings.ToUpper(scanner.Text())
	if inputTipe != "A" && inputTipe != "B" && inputTipe != "C" && inputTipe != "D" {
		fmt.Print("Input tidak valid")
		return
	}

	fmt.Print("Input jumlah deret: ")
	scanner.Scan()
	inputJumlahDeret := scanner.Text()
	jumlahDeret, err := strconv.Atoi(inputJumlahDeret)
	if err != nil {
		fmt.Print("Input tidak valid")
		return
	}

	var counter int
	switch inputTipe {
	case "A":
		for index := 0; index < jumlahDeret; index++ {
			counter += 1
			fmt.Printf("%d ", counter)
		}
	case "B":
		for index := 0; index < jumlahDeret; index++ {
			counter += 3
			fmt.Printf("%d ", counter)
		}
	case "C":
		counter = 7
		for index := 0; index < jumlahDeret; index++ {
			counter += 1
			fmt.Printf("%d ", counter)
		}
	case "D":
		counter = 1
		ganjil := 1
		for index := 0; index < jumlahDeret; index++ {
			if index != 0 {
				ganjil += 2
				counter += ganjil
			}
			fmt.Printf("%d ", counter)
		}
	}
}
