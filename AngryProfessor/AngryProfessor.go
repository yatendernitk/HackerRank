package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var totalNum int
	fmt.Scan(&totalNum)
	reader := bufio.NewReader(os.Stdin)
	totalPresent := 0
	studentTime := 0
	for i:=0 ; i< totalNum ; i++ {
		totalPresent = 0
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)

		s1 := strings.Split(s, " ")
		totalStudents,_ := strconv.Atoi(s1[0])
		thresholdStudents,_ := strconv.Atoi(s1[1])

		s2, _ := reader.ReadString('\n')
		s2 = strings.TrimSpace(s2)
		s21 := strings.Split(s2, " ")

		if len(s21) != totalStudents {
			fmt.Println("erorororor")
		}else{
			for j:= 0 ; j< totalStudents ; j++ {
				studentTime,_ = strconv.Atoi(s21[j])
				if studentTime > 0 {
					// do nothing
				}else{
					totalPresent++
				}
			}
		}
		// finish each iteration
		if totalPresent >= thresholdStudents{
			fmt.Println("NO")
		}else{
			fmt.Println("YES")
		}

	}

}
