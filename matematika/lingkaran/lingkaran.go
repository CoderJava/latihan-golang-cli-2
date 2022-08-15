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

	fmt.Print("Input jari-jari lingkaran: ")
	scanner.Scan()
	strInputJariJari := scanner.Text()
	inputJariJari, err := strconv.ParseFloat(strInputJariJari, 64)
	if err != nil {
		fmt.Println("Input jari-jari tidak valid.")
		return
	}

	luas := 3.14 * math.Pow(inputJariJari, 2)
	fmt.Printf("Luas lingkaran: %f", luas)
}
