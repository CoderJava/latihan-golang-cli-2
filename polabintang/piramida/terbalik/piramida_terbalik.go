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
	inputTinggiSegitga, err := strconv.Atoi(strInputTinggiSegitiga)
	if err != nil {
		fmt.Println("Input tinggi segitiga tidak valid.")
		return
	}

	for x := 0; x < inputTinggiSegitga; x++ {
		for y := 0; y < inputTinggiSegitga; y++ {
			if y < x {
				fmt.Print(" ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}

}

/*
tinggi: 4
* * * *
 * * *
  * *
   *

tinggi: 5
* * * * *
 * * * *
  * * *
   * *
    *
*/
