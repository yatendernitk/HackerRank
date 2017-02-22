package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	var totalCase int
	fmt.Scan(&totalCase)
	reader := bufio.NewReader(os.Stdin)

	for i:=0; i< totalCase ; i++{
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		inputItem := strings.Split(text, " ")

		totalItems, _ := strconv.Atoi(inputItem[0])
		maxCapacity,_ := strconv.Atoi(inputItem[1])



		capacityItemArr := make([]int,totalItems)

		for k:=0 ; k < totalItems ; k++ {
			fmt.Scan(&capacityItemArr[k])
		}

		sort.Ints(capacityItemArr)

		cap := 0
		lol := 0

		for j:=0 ; j < totalItems ; j++ {
			if cap+capacityItemArr[j] <= maxCapacity {
				cap += capacityItemArr[j]
				lol++
			}else{
				break;
			}

		}
		fmt.Println(lol)

	}
}
