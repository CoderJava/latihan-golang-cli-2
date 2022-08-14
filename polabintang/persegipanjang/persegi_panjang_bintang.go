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

	fmt.Print("Input tinggi persegi: ")
	scanner.Scan()
	strInputTinggiPersegi := scanner.Text()
	inputTinggiPersegi, err := strconv.Atoi(strInputTinggiPersegi)
	if err != nil {
		fmt.Println("Input tinggi persegi tidak valid.")
		return
	}

	fmt.Print("Input lebar persegi: ")
	scanner.Scan()
	strInputLebarPersegi := scanner.Text()
	inputLebarPersegi, err := strconv.Atoi(strInputLebarPersegi)
	if err != nil {
		fmt.Println("Input lebar persegi tidak valid.")
		return
	}

	for x := 0; x < inputTinggiPersegi; x++ {
		var str string
		for y := 0; y < inputLebarPersegi; y++ {
			str += "* "
		}
		str = strings.TrimRight(str, " ")
		fmt.Println(str)
	}
}
