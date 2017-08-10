package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var totalNum int
	fmt.Println("enter square matrix row size")
	fmt.Scan(&totalNum)

	values := [][]string{}
	reader := bufio.NewReader(os.Stdin)

	numstr := make([]string, totalNum)
	fmt.Println("enter  matrix row")
	// These are the first two rows.

	for i := 0; i < totalNum; i++ {

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		numstr = strings.Split(text, " ")
		values = append(values, numstr)

	}

	for j := 0; j < totalNum; j++ {
		jlen := len(values[j])

		if isEven(j) {
			fmt.Print(values[j])
		} else {
			for k := 0; k < jlen; k++ {
				fmt.Print(values[j][jlen-k-1] + " ")
			}
		}
		fmt.Println("")
	}
}

func isEven(number int) bool {
	return number%2 == 0
}

