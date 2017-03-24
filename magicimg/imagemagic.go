package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	//"os"
)

func main() {
	// Initialize the graphic context on an RGBA image
	//fImg1, _ := os.Open("src.jpg")
	//defer fImg1.Close()
	//img1, _, _ := image.Decode(fImg1)

	//var m = image.NewRGBA(0,0,600)

	dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(1)

	// Draw a closed shape
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)

	gc.Close()
	gc.FillStroke()


	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)
}