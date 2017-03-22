package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	str := strings.TrimSpace(text)
	x := ""
	y := 0
	for i := 0; i < len(str); i++ {
		x = str[i : i+1]
		if isUpperCase(x) {
			//fmt.Println("")
			y = y + 1

		} else {

			//fmt.Print(x)
		}

	}
	y = y + 1
	fmt.Println(y)

}

func isUpperCase(x string) bool {
	if x == strings.ToUpper(x) {
		return true
	} else {
		return false
	}
}
