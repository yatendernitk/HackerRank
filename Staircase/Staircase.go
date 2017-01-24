package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n')
	str := strings.TrimSpace(text)
	num,err := strconv.Atoi(str)
	var spaceNum int
	if err != nil {
		fmt.Println(err)
	}

	for i:=1 ; i <= num ; i++ {
		spaceNum = num-i;
		for j:= 1 ; j<= num ; j++ {
			if j<=spaceNum{
				fmt.Print(" ")
			}else{
				fmt.Print("#")
			}
		}
		fmt.Println("")

	}
}
