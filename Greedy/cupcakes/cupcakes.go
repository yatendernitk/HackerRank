package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := readdata()
	fmt.Println(do_operation(nums))
}

func do_operation(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		pownum := int(math.Pow(2, float64(i))) * nums[i]
		sum += pownum
	}
	return sum
}

func stringToNumArray(strArr []string) []int {
	inArr := make([]int, len(strArr))
	for i := range strArr {
		inArr[i], _ = strconv.Atoi(strArr[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inArr)))
	return inArr
}

func readdata() []int {
	var totalCase int
	fmt.Scan(&totalCase)
	reader := bufio.NewReader(os.Stdin)
	text2, _ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	numArr := strings.Split(text2, " ")
	nums := stringToNumArray(numArr)
	return nums
}
