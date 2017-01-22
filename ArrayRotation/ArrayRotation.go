package main

import ("fmt"
	"bufio"
	"os"
	"strings"
	"errors"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	var totalNum int
	fmt.Scan(&totalNum)

	text, _ := reader.ReadString('\n')

	s := strings.TrimSpace(text)
	result := strings.Split(s, " ")


	if len(result) != totalNum{
		fmt.Print("I am here")
		  errors.New("can't work with 42")
	}else{
		for i := len(result) - 1; i >= 0; i-- {
			fmt.Print(result[i])
			if i > 0 {
				fmt.Print(" ")
			}
		}
	}
	// Display all elements.

}