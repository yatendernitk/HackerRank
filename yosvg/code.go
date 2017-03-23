package main

import (
	"github.com/ajstarks/svgo"
	"os"
)

func main(){
	width := 512
	height := 512

	canvas := svg.New(os.Stdout)
	canvas.Start(width,height)
	canvas.Title("canvas fade")
	canvas.Image(0,0,400,400,"src.jpg","0.50")
	canvas.End()

}
