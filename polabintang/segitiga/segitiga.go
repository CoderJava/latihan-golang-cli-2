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

	fmt.Print("Input tinggi segitiga: ")
	scanner.Scan()
	strInputTinggSegitiga := scanner.Text()
	inputTinggiSegitiga, err := strconv.Atoi(strInputTinggSegitiga)
	if err != nil {
		fmt.Println("Input tinggi segitiga tidak valid.")
		return
	}

	for x := 0; x < inputTinggiSegitiga; x++ {
		var str string
		for y := 0; y <= x; y++ {
			str += "* "
		}
		str = strings.TrimRight(str, " ")
		fmt.Println(str)
	}
}
