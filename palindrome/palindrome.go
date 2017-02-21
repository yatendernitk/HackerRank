package main

import (
	//"bufio"
	//"os"
	//"strings"
	"fmt"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//text,_ := reader.ReadString('\n')
	//text = strings.TrimSpace(text)
	result := checkPalindrome("abccba")
	fmt.Println(result)
}

func checkPalindrome(inputString string) bool {
	l := len(inputString)
	for i:=0; i< l ; i++ {
		if inputString[i:i+1] != inputString[l-1-i:l-i] {
			return false
		}
	}
	return true
}

