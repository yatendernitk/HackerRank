package main

import (
	"github.com/ajstarks/svgo"
	"os"
	_ "image"
	_ "fmt"
)

func main(){
	width := 512
	height := 512

	canvas := svg.New(os.Stdout)
	canvas.Start(width,height)
	canvas.Image(0,0,512,512,"src.jpg","0.50")
	canvas.End()
	//k := canvas
}
