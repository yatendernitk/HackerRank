package main

import (
	"io/ioutil"
	"fmt"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	dat,err := ioutil.ReadFile("sahiba.mp3")
	check(err)
	fmt.Print(dat)
}
