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
	//var q int
	//var qNum string
	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimSpace(firstLine)
	l1 := strings.Split(firstLine, " ")


	secondLine, _ := reader.ReadString('\n')
	secondLine = strings.TrimSpace(secondLine)
	l2 := strings.Split(secondLine, " ")

	lenOfArray,eLenOfArray := strconv.Atoi(l1[0])
	rotation,eRotation := strconv.Atoi(l1[1])
	queries ,eQueries := strconv.Atoi(l1[2])
	finalArray := make([]string, lenOfArray)

	if eLenOfArray != nil || eRotation != nil || eQueries != nil ||
		lenOfArray < 1 || lenOfArray > 100000 ||
		rotation < 1 || rotation > 100000 ||
		queries < 1 || queries > 500 ||
		len(l2) != lenOfArray{
		fmt.Print("erorr")
	}else{
		k := 0
		for i:= len(l2)-rotation ; i < len(l2) ;  i++ {
			//lets assume we want to rotate array twice
			//length of array() - n
			// 3 2 1 is there so we will print after first rotation 1 3 2 and 2 1 3 after second rotation
			//fmt.Print(l2[i], " ");
			//fmt.Println(i,":", l2[i])
			finalArray[k] = l2[i]

			//if err!= nil {
			//	fmt.Print("")
			//}
			//fmt.Println("get:", finalArray)
			k++
		}
		for i :=0 ; i < len(l2)-rotation ; i++ {
			//fmt.Println(i,":", l2[i])
			finalArray[k] = l2[i]
			k++
		}

		//for i:= range l2 {
		//	fmt.Print(finalArray[i]," ")
		//}

		//fmt.Println("get Final :", finalArray)
		//fmt.Println("")
		//for i:=0 ; i < queries ; i++ {
		//	fmt.Scan(&q)
		//	fmt.Println(finalArray[q], " ")
		//}
		queryArray := make([]int, queries)
		for i:=0; i< queries ; i++ {
			//qNum, _ = reader.ReadString('\n')
			//fmt.Println(qNum, ":", i)
			//queryArray[i],_ = strconv.Atoi(qNum)
			fmt.Scanf("%d\n",&queryArray[i])
			//fmt.Println(queryArray[i])
			//fmt.Println(queryArray[i]," ",i)
		}

				//fmt.Println(queryArray)

		for i:=range queryArray{
			fmt.Println(finalArray[queryArray[i]])
		}

		// my answer is correct but i don't know somehow it's not working in hackerrank

	}
	//
	//var n int
	//fmt.Scan(&n)


}
