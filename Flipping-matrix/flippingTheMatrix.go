package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"sort"
)

func main() {
	var totalInput int
	fmt.Scan(&totalInput)
	if totalInput <1 || totalInput > 16{
		panic("eroror")
	}
	reader := bufio.NewReader(os.Stdin)
	for s:=0; s< totalInput; s++ {
		fmt.Println("s - ",s)
		var n int
		fmt.Scan(&n)
		if n < 1 || n > 128 {
			panic("erororr")
		}
		totalNum := 2*n

		a := make([][]int, totalNum)
		e := make([]int, totalNum * totalNum)


		numstr := make([]string, totalNum)


		for i := range a {
			a[i] = e[i*totalNum:(i+1)*totalNum]
			fmt.Println(i)
		}


		for i:=0 ; i < totalNum ; i++ {
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			numstr = strings.Split(text, " ")
			for j:=0 ; j < totalNum ; j++ {
				a[i][j],_ = strconv.Atoi(numstr[j])
				if a[i][j] < 0 || a[i][j] > 4096{
					panic("error")
				}
			}
		}
		sum := 0
		for i := 0; i < n ; i++ {
			for  j := 0; j < n ; j++{
				sum += max(a[i][j], a[i][totalNum - j - 1], a[totalNum - i - 1][j], a[totalNum - i - 1][totalNum - j - 1]);
				fmt.Println(i,j)
			}
		}

		fmt.Println(sum)
	}

}

func max(a,b,c,d int) int {
	k := make([]int , 4)
	k[0] = a
	k[1] = b
	k[2] = c
	k[3] = d
	sort.Ints(k)
	return k[3]
}
