package main

import (
	"fmt"
	"sort"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text,_ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	inArr := strings.Split(text," ")
	totalFlowers,_ := strconv.Atoi(inArr[0])
	numPeople,_ := strconv.Atoi(inArr[1])

	flowers := make([]int,totalFlowers)

	textLine2,_ := reader.ReadString('\n')
	textLine2 = strings.TrimSpace(textLine2)
	floArr := strings.Split(textLine2, " ")

	for i:=0 ; i< totalFlowers ; i++ {
		flowers[i],_ =strconv.Atoi(floArr[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(flowers)))
	//fmt.Println(flowers)
	result := calculateCost(flowers,numPeople)
	fmt.Println(result)


}


func calculateCost(flowers []int,numPeople int) int {
	//flowers are unsorted
	totalFlowers := len(flowers)
	cost:=0
	//fmt.Println(totalFlowers,numPeople)

	//fmt.Println(totalFlowers/numPeople)

	if totalFlowers >= numPeople {
		flowerIndex :=0

		for i:= 0;flowerIndex < totalFlowers; i++ {
			for j:=0; j< numPeople; j++ {
				if flowerIndex >= totalFlowers {
					break;
				}else{
					//fmt.Println(flowerIndex)
					//fmt.Println(cost,flowers[flowerIndex],i)
					cost += (i+1)* flowers[flowerIndex]
					flowerIndex++
					//fmt.Println(flowerIndex)
					//fmt.Println(cost)
				}

			}

		}

		//fmt.Println(flowerIndex)
		//
		//k := 1;
		//
		//for j:=flowerIndex ; j < totalFlowers ; j++ {
		//
		//	fmt.Println(flowerIndex,totalFlowers,k)
		//	fmt.Println((k+1)*flowers[flowerIndex])
		//	cost += (k+1)*flowers[flowerIndex]
		//	k++
		//	flowerIndex++
		//
		//}

	}else{
		for i:= 0;i < totalFlowers; i++ {
			cost += flowers[i]
		}
	}

	return cost
}