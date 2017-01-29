package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
	"fmt"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text,_ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	inArr := strings.Split(text," ")

	text2,_ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	numArr := strings.Split(text2," ")
	intNumArr := stringToNumArray(numArr)
	diffNum := 0
	arrSize,_ := strconv.Atoi(inArr[0])
	diff ,_ := strconv.ParseFloat(inArr[1],64)
	j:= arrSize-1
	if arrSize != len(numArr) {
		panic("errrorrr")
	}else{

		for i:= range intNumArr{

			//for j:= range numArr{
			//
			//	if(math.Abs(intNumArr[i]-intNumArr[j]) == diff){
			//		diffNum++
			//	}
			//}
			//fmt.Println(i)
			j = arrSize-1
			for  ;math.Abs(intNumArr[i]-intNumArr[j]) > diff && j > 0 ; {
				//fmt.Println("up")
				//fmt.Println(intNumArr[i],intNumArr[j])
				j--
			}
			if math.Abs(intNumArr[i]-intNumArr[j]) == diff {
				//fmt.Println("down")
				//fmt.Println(intNumArr[i],intNumArr[j])
				diffNum++
			}
		}

		}
		fmt.Println(diffNum)
	}






//while (i < a.size()) {
//while(a[i] + b[j] > X && j > 0) j--;
//if (a[i] + b[j] == X) writeAnswer(i, j);
//i++;
//}

func stringToNumArray(strArr []string) []float64 {
	inArr := make([]float64, len(strArr))
	for i:=range strArr {
		inArr[i],_ = strconv.ParseFloat(strArr[i],64)
	}
	sort.Float64s(inArr)
	//fmt.Println(inArr)
	return inArr
}


