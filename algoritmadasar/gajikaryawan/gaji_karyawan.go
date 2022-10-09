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
	fmt.Println("## Program Menghitung Gaji Karyawan ##")
	fmt.Println("======================================")

	// Input nama karyawan
	fmt.Print("Nama karyawan: ")
	scanner.Scan()
	strInputNamaKaryawan := scanner.Text()

	// Input golongan
	fmt.Print("Golongan: ")
	scanner.Scan()
	strInputGolongan := scanner.Text()

	//  Input jumlah jam kerja
	fmt.Print("Jumlah jam kerja: ")
	scanner.Scan()
	strInputJumlahJamKerja := scanner.Text()
	inputJumlahJamKerja, err := strconv.Atoi(strInputJumlahJamKerja)
	if err != nil {
		fmt.Print("Input jumlah jam kerja tidak valid")
		return
	}

	var upah int
	switch strings.ToUpper(strInputGolongan) {
	case "A":
		upah = 5000 * inputJumlahJamKerja
	case "B":
		upah = 7000 * inputJumlahJamKerja
	case "C":
		upah = 8000 * inputJumlahJamKerja
	case "D":
		upah = 10000 * inputJumlahJamKerja
	default:
		fmt.Print("Input golongan tidak valid")
		return
	}

	if inputJumlahJamKerja > 48 {
		uangLembur := (inputJumlahJamKerja - 48) * 4000
		upah += uangLembur
	}

	fmt.Printf("%s menerima upah Rp.%d per minggu", strInputNamaKaryawan, upah)
}
