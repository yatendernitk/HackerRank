package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := readdata1()
	dollars := nums[1]
	toysPrice := readdata()
	totalToys := 0
	for _, num := range toysPrice {
		if dollars-num > 0 {
			dollars -= num
			totalToys++
		}
	}
	fmt.Println(totalToys)
}
func readdata() []int {
	reader := bufio.NewReader(os.Stdin)
	text2, _ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	numArr := strings.Split(text2, " ")
	nums := stringToNumArray(numArr)
	return nums
}

func readdata1() []int {

	reader := bufio.NewReader(os.Stdin)
	text2, _ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	numArr := strings.Split(text2, " ")
	nums := stringToNumArray1(numArr)
	return nums
}

func stringToNumArray(strArr []string) []int {
	inArr := make([]int, len(strArr))
	for i := range strArr {
		inArr[i], _ = strconv.Atoi(strArr[i])
	}
	sort.Sort(sort.IntSlice(inArr))
	return inArr
}

func stringToNumArray1(strArr []string) []int {
	inArr := make([]int, len(strArr))
	for i := range strArr {
		inArr[i], _ = strconv.Atoi(strArr[i])
	}
	return inArr
}
