package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pilih tipe [1/2/3/4]: ")
	scanner.Scan()
	inputTipe := scanner.Text()
	if inputTipe != "1" && inputTipe != "2" && inputTipe != "3" && inputTipe != "4" {
		fmt.Print("Input tidak valid")
		return
	}

	switch inputTipe {
	case "1":
		fmt.Print("Jumlah deret yang diinginkan: ")
		scanner.Scan()
		strInputJumlahDeret := scanner.Text()
		inputJumlahDeret, err := strconv.Atoi(strInputJumlahDeret)
		if err != nil {
			fmt.Print("Input jumlah deret tidak valid")
			return
		}
		counter := 0
		for x := 1; x <= inputJumlahDeret; x++ {
			counter += 2
			fmt.Print(counter)
			if x != inputJumlahDeret {
				fmt.Print(" ")
			}
		}
	case "2":
		fmt.Print("Jumlah deret yang diinginkan: ")
		scanner.Scan()
		strInputJumlahDeret := scanner.Text()
		inputJumlahDeret, err := strconv.Atoi(strInputJumlahDeret)
		if err != nil {
			fmt.Print("Input jumlah deret tidak valid")
			return
		}
		counter := -1
		for x := 1; x <= inputJumlahDeret; x++ {
			counter += 2
			fmt.Print(counter)
			if x != inputJumlahDeret {
				fmt.Print(" ")
			}
		}
	case "3":
		fmt.Print("Input angka awal: ")
		scanner.Scan()
		strInputAngkaAwal := scanner.Text()
		inputAngkaAwal, err := strconv.Atoi(strInputAngkaAwal)
		if err != nil {
			fmt.Print("Input angka awal tidak valid")
			return
		}

		fmt.Print("Input angka akhir: ")
		scanner.Scan()
		strInputAngkaAkhir := scanner.Text()
		inputAngkaAkhir, err := strconv.Atoi(strInputAngkaAkhir)
		if err != nil {
			fmt.Print("Input angka akhir tidak valid")
			return
		}
		if inputAngkaAkhir%2 != 0 {
			inputAngkaAkhir -= 1
		}

		counter := inputAngkaAwal
		if inputAngkaAwal%2 != 0 {
			counter = inputAngkaAwal + 1
		}
		for counter <= inputAngkaAkhir {
			fmt.Print(counter)
			if counter != inputAngkaAkhir {
				fmt.Print(" ")
			}
			counter += 2
		}
	case "4":
		fmt.Print("Input angka awal: ")
		scanner.Scan()
		strInputAngkaAwal := scanner.Text()
		inputAngkaAwal, err := strconv.Atoi(strInputAngkaAwal)
		if err != nil {
			fmt.Print("Input angka awal tidak valid")
			return
		}

		fmt.Print("Input angka akhir: ")
		scanner.Scan()
		strInputAngkaAkhir := scanner.Text()
		inputAngkaAkhir, err := strconv.Atoi(strInputAngkaAkhir)
		if err != nil {
			fmt.Print("Input angka akhir tidak valid")
			return
		}
		if inputAngkaAkhir%2 == 0 {
			inputAngkaAkhir -= 1
		}

		counter := inputAngkaAwal
		if inputAngkaAwal%2 == 0 {
			counter += 1
		}
		for counter <= inputAngkaAkhir {
			fmt.Print(counter)
			if counter != inputAngkaAkhir {
				fmt.Print(" ")
			}
			counter += 2
		}
	}
}
