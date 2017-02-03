package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	m := make(map[string]int)
	cArr := strings.Split(text,"")
	if len(cArr) > 100000 {
		panic("errorr")
	}else{
		for i:= range cArr {
			if isKeyExists(m,cArr[i] ) {
				//delete the key and increment the value
				//fmt.Println("key exists")
				deleteKey(m,cArr[i])
			}else{
				// add the key in map
				//fmt.Println("key not exists")
				addKey(m,cArr[i])
			}
		}
		//now only 1 or 0 key should be there
		//fmt.Println(len(m))
		if len(m) >1 {
			fmt.Println("NO")
		}else{
			fmt.Println("YES")
		}
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
	//fmt.Println(key)
	delete(m, key)
}
func addKey(m map[string]int, key string){
	//fmt.Println("addKey ",key)
	m[key] = 1
}