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
	var totalCase int
	fmt.Scan(&totalCase)
	reader := bufio.NewReader(os.Stdin)
	text2, _ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	numArr := strings.Split(text2, " ")
	nums := stringToNumArray(numArr)
	diff := 9999999.0

	for i := 0; i < totalCase-1; i++ {
		if math.Abs(nums[i]-nums[i+1]) < diff {
			diff = math.Abs(nums[i] - nums[i+1])
		}
	}
	fmt.Println(diff)
}

func stringToNumArray(strArr []string) []float64 {
	inArr := make([]float64, len(strArr))
	for i := range strArr {
		inArr[i], _ = strconv.ParseFloat(strArr[i], 64)
	}
	sort.Float64s(inArr)
	//fmt.Println(inArr)
	return inArr
}
