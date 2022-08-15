package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input jumlah element Array: ")
	scanner.Scan()
	strInputJumlahElement := scanner.Text()
	inputJumlahElement, err := strconv.Atoi(strInputJumlahElement)
	if err != nil {
		fmt.Println("Input jumlah element array tidak valid.")
		return
	}

	fmt.Printf("Input %d angka (dipisah dengan enter):", inputJumlahElement)
	fmt.Println()
	elements := make([]int, inputJumlahElement)
	isValid := true
	for x := 0; x < inputJumlahElement; x++ {
		scanner.Scan()
		strInputElement := scanner.Text()
		inputElement, err := strconv.Atoi(strInputElement)
		if err != nil {
			fmt.Println("Input angka tidak valid.")
			isValid = false
			break
		}
		elements[x] = inputElement
	}
	if !isValid {
		return
	}
	fmt.Println()

	fmt.Print("Input angka yang akan dicari: ")
	scanner.Scan()
	strInputAngkaDicari := scanner.Text()
	inputAngkaDicari, err := strconv.Atoi(strInputAngkaDicari)
	if err != nil {
		fmt.Println("Input angka yang akan dicari tidak valid.")
		return
	}

	// Sequential searching
	/* fmt.Println(elements)
	isFounded := false
	for index, element := range elements {
		if element == inputAngkaDicari {
			fmt.Printf("Angka ditemukan pada index ke-%d", index)
			isFounded = true
			break
		}
	}
	if !isFounded {
		fmt.Println("Angka tidak ditemukan")
	} */

	// Binary searching
	startIndex := 0
	stopIndex := len(elements) - 1
	centerIndexOld := -1
	counterCenterIndexOld := 0
	sort.Ints(elements)
	fmt.Println(elements)
	for {
		centerIndex := (startIndex + stopIndex) / 2
		valueOfCenterIndex := elements[centerIndex]
		if valueOfCenterIndex == inputAngkaDicari {
			fmt.Printf("Angka ditemukan pada index ke-%d", centerIndex)
			break
		} else if valueOfCenterIndex < inputAngkaDicari {
			startIndex = centerIndex + 1
		} else if valueOfCenterIndex > inputAngkaDicari {
			stopIndex = centerIndex - 1
		}
		if centerIndex == centerIndexOld {
			counterCenterIndexOld++
			if counterCenterIndexOld == 3 {
				fmt.Println("Angka tidak ditemukan")
				break
			}
		}
		centerIndexOld = centerIndex
	}
}
