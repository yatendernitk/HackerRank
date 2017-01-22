package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
	"errors"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	s := strings.TrimSpace(text)
	//s = "09:05:45PM"
	inputTime := strings.Split(s, ":")



	firstS := inputTime[0]

	firstNum, err := strconv.Atoi(firstS)
	secondNum, err2 := strconv.Atoi(inputTime[1])

	if(firstNum>12 || firstNum < 01 || secondNum > 59 || secondNum < 0){
		errors.New("wrong input")
		os.Exit(1)
	}

	if err != nil && err2 != nil {
		os.Exit(1)
	}


	lastS := inputTime[2]
	thirdNum,err3 := strconv.Atoi(lastS[:2])

	if(thirdNum > 59 || thirdNum < 0){
		errors.New("wrong input")
		os.Exit(1)
	}



	if err3 != nil {
		fmt.Println("error")
	}
	timeFactor := lastS[2:]
	if(thirdNum == 0){

	}
	if timeFactor == "AM" {
		if(firstNum == 12){
			fmt.Println("00:"+inputTime[1]+":"+lastS[:2])
		}else{
			fmt.Println(inputTime[0]+":"+inputTime[1]+":"+lastS[:2])
		}

	}else{

		if(firstNum == 12){
			fmt.Println("12:"+inputTime[1]+":"+lastS[:2])
		}else{
			finalNum := (firstNum + 12)%24

			if finalNum == 0 {
				fNum := "00"
				fmt.Println(fNum + ":" + inputTime[1] + ":" +lastS[:2])
			}else{
				fmt.Println(strconv.Itoa(finalNum)+":"+inputTime[1]+":"+lastS[:2])

			}
		}




	}

}
