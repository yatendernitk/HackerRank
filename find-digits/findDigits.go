package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	var totalNum int
	fmt.Scan(&totalNum)
	reader := bufio.NewReader(os.Stdin)
	numDiv := 0
	numInInt := 0

	var splittedString []string
	for i:=0 ; i< totalNum ; i++ {
		numDiv = 0
		num,_ := reader.ReadString('\n')
		num = strings.TrimSpace(num)
		finalNum,_ := strconv.Atoi(num)

		splittedString = strings.Split(num, "")
		//fmt.Println(splittedString)
		for j:= range splittedString {
			numInInt,_ = strconv.Atoi(splittedString[j])
			//fmt.Println(numInInt)
			if numInInt != 0 {
				if finalNum%numInInt>0  {

					//fmt.Println("we r heere")
				}else{
					numDiv++
				}
			}
		}
		fmt.Println(numDiv)
	}
}
