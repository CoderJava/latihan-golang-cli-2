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
	fmt.Print("Input besar persegi: ")
	scanner.Scan()
	strInputAngka := scanner.Text()
	inputAngka, err := strconv.Atoi(strInputAngka)
	if err != nil {
		fmt.Println("Input angka tidak valid")
		return
	}

	for x := 0; x < inputAngka; x++ {
		var str string
		for y := 0; y < inputAngka; y++ {
			str += "* "
		}
		str = strings.TrimRight(str, " ")
		fmt.Println(str)
	}
}
