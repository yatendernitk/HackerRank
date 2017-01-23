package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	sum := 0
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	str := strings.TrimSpace(text)
	result := strings.Split(str, " ")
	var min int
	var max int
	var err error

	for i := 0; i < 5; i++ {
		num,err2 := strconv.Atoi(result[i])
		if i == 0{
			min,err = strconv.Atoi(result[i])
		}
		if err != nil || err2 != nil {

		}
		min = findMin(min,num)
		max = findMax(max,num)
		sum += num

	}
	//fmt.Println(sum)
	fmt.Print(sum-min)
	fmt.Print(" ")
	fmt.Print(sum-max)
}

func findMin(min ,num int) int {
	if min < num{
		return min
	}else{
		return num
	}
}

func findMax(max ,num int) int {
	if max > num{
		return max
	}else{
		return num
	}
}
