package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	var totalNum int
	fmt.Scan(&totalNum)
	reader := bufio.NewReader(os.Stdin)
	text:=""

	for i:=0 ; i< totalNum ; i++ {
		primes := [10]string{"h", "a", "c", "k", "e", "r","r","a","n","k"}
		text,_ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
		s := strings.Split(text,"")
		k:=0
		sum:=0
		for j:= range s {
			if(s[j]==primes[k]){
				sum += 1
				fmt.Println(k)
				if(k==9){
					break;
				}
				k++
			}
		}

		if sum != 10 {
			fmt.Println("NO")
		}else{
			fmt.Println("YES")
		}
	}



}

