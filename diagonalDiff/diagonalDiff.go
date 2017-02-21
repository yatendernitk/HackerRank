package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func main(){
	var totalNum int
	fmt.Scan(&totalNum)
	middleNum := totalNum/2
	fmt.Println(middleNum)
	reader := bufio.NewReader(os.Stdin)

	numstr := make([]string, totalNum)

	a := make([][]int, totalNum)
	e := make([]int, totalNum * totalNum)

	d1 := 0
	d2 := 0

	for i := range a {
		a[i] = e[i*totalNum:(i+1)*totalNum]
	}

	for i:=0 ; i < totalNum ; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		numstr = strings.Split(text, " ")
		for j:=0 ; j < totalNum ; j++ {
			a[i][j],_ = strconv.Atoi(numstr[j])
			//fmt.Println(i,j)

			if i==j {
				d1 += a[i][j]
			}
		}
	}
	k := totalNum-1
	for j:=0; j<totalNum; j++{
		if k>=0{
			d2=d2+a[j][k];
			k--;
		}
	}
	fmt.Println(d1,d2)
	fmt.Println(math.Abs(float64(d1)-float64(d2)))

}
