package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"errors"
	"strconv"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	var totalNum float64
	fmt.Scan(&totalNum)

	text, _ := reader.ReadString('\n')
	str := strings.TrimSpace(text)
	result := strings.Split(str, " ")

	var positive float64
	var zero float64
	var negative float64


	if float64(len(result))  != totalNum{
		errors.New("can't work with 42")
	}else{
		for i := len(result) - 1; i >= 0; i-- {
			num,err := strconv.Atoi(result[i])

			if err != nil {
				errors.New("can't work with 42")
			}

			if  num < 0 {
				negative = negative +1;
			} else if num == 0 {
				zero = zero + 1;
			} else {
				positive = positive +1;
			}
		}
		totalNegNum := negative/totalNum
		totalPosNum := positive/totalNum
		totalZeroNum := zero/totalNum



		fmt.Println(totalPosNum);
		fmt.Println(totalNegNum);
		fmt.Println(totalZeroNum);

	}
}