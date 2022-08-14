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
	strInputTinggiSegitiga := scanner.Text()
	inputTinggiSegitiga, err := strconv.Atoi(strInputTinggiSegitiga)
	if err != nil {
		fmt.Println("Input tinggi segitiga tidak valid.")
		return
	}

	for x := inputTinggiSegitiga; x >= 1; x-- {
		var str string
		for y := 1; y <= x; y++ {
			str += "* "
		}
		str = strings.TrimRight(str, " ")
		fmt.Println(str)
	}
}
