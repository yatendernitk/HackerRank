package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text1, _ := reader.ReadString('\n')
	s1 := strings.TrimSpace(text1)
	aliceList := strings.Split(s1, " ")

	text2, _ := reader.ReadString('\n')
	s2 := strings.TrimSpace(text2)
	bobList := strings.Split(s2, " ")


	aliceS := 0
	bobS:= 0

	if len(aliceList) != len(bobList){
		fmt.Println("errror")
	}else{
		for i := range bobList {

			aliceNum, err := strconv.Atoi(aliceList[i])
			bobNum, err2 := strconv.Atoi(bobList[i])

			if err != nil && err2 != nil {
				os.Exit(1)
			}

			if aliceNum > bobNum {
				//fmt.Println("with alics ")
				aliceS = aliceS + 1
			}else if aliceNum < bobNum {
				//fmt.Println("with  bob")
				bobS = bobS + 1
			}else {
				//fmt.Println("in else")
			}
		}


			fmt.Print(aliceS," ")



			fmt.Print(bobS)


	}
}
