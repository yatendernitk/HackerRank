package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var totalNum int
	fmt.Scan(&totalNum)
	m := make(map[string]int)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	numList := strings.Split(text, " ")
	numPairs := 0
	if totalNum != len(numList){
		fmt.Print("error")
	}else{
		for i:= range numList  {
			if isKeyExists(m,numList[i] ) {
				//delete the key and increment the value
				deleteKey(m,numList[i])
				//fmt.Println("exists")
				numPairs += 1
			}else{
				// add the key in map
				addKey(m,numList[i])
				//fmt.Println("not exists")
			}
		}

		fmt.Println(numPairs)

	}
}

func isKeyExists(m map[string]int, key string) bool {
	//fmt.Println("isKeyExists ",key)
	_,ok := m[key]
	//fmt.Println(ok)
	return ok
}

func deleteKey(m map[string]int, key string){
	//fmt.Println("deleteKey ",key)
	delete(m, key)
}
func addKey(m map[string]int, key string){
	//fmt.Println("addKey ",key)
	m[key] = 1
}