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
	inputTinggiSegitiga, err := strconv.Atoi(strInputTinggiSegitiga)
	if err != nil {
		fmt.Println("Input tinggi segitiga tidak valid")
		return
	}

	for x := 0; x < inputTinggiSegitiga; x++ {
		limitSpace := inputTinggiSegitiga - x
		isPrintStar := true
		for y := 1; y <= inputTinggiSegitiga+x+1; y++ {
			if y <= limitSpace {
				fmt.Print(" ")
			} else {
				if isPrintStar {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
				isPrintStar = !isPrintStar
			}
		}
		fmt.Println()
	}
}

/*
tinggi: 4

    *
   * *
  * * *
 * * * *

tinggi: 5
     *          6
    * *			5 7
   * * *		4 6 8
  * * * *		3 5 7 9
 * * * * *		2 4 6 8 10
*/
