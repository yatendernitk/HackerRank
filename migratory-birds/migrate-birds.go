package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var totalNum int
	fmt.Scan(&totalNum)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	numList := strings.Split(text, " ")
	m := make(map[int]int)
	number:=0
	val := 0


	for i:=range numList {
		number,_ = strconv.Atoi(numList[i])
		if isKeyExists(m,number ) {
			//delete the key and increment the value
			val = getKeyVal(m,number)
			val++
			updateKey(m,number,val)
			//fmt.Println("exists")
		}else{
			// add the key in map
			addKey(m,number)
			//fmt.Println("not exists")
		}
	}

	number=0
	val = 0
	maxFreq := 0
	minNum := 200000
	for key, value := range m {
		if(value==maxFreq){
			maxFreq = value
			if(minNum>key){
				minNum = key
			}

		}else if(value>maxFreq){
			maxFreq = value
			minNum = key
		}
	}

	fmt.Println(minNum)

}


func isKeyExists(m map[int]int, key int) bool {
	//fmt.Println("isKeyExists ",key)
	_,ok := m[key]
	//fmt.Println(ok)
	return ok
}

func addKey(m map[int]int, key int){
	m[key] = 1
}

func updateKey(m map[int]int, key int, val int){
	m[key] = val
}

func getKeyVal(m map[int]int, key int) int{
	return m[key]
}
