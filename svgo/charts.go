package main

import (
	"os"
	"image/draw"
	"image"
	"image/jpeg"
	"image/color"
	//"strconv"
	"fmt"
)

	//draw.Draw(mask, mask.Bounds(), image.Transparent, image.ZP, draw.Over)

func main() {
	//runtime.GOMAXPROCS(4)
	//start := time.Now()
	//for i:=0; i< 500 ; i++ {
	//	fmt.Println(i)
		drawImg()
	//}
	//nm := ""
	//fmt.Scanf("%s",&nm)
	//elapsed := time.Since(start)
	//log.Printf("Binomial took %s", elapsed)
}


func drawImg() {
	//wg.Add(1)
	//go func(){

	fImg1, _ := os.Open("src.jpg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)
	fmt.Println(img1.Bounds())

	fImg2, _ := os.Open("logo.jpg")
	defer fImg2.Close()

	fImg3, _ := os.Open("Black.jpg")
	defer fImg3.Close()
	img3,_,_ := image.Decode(fImg3)
	img2, _, _ := image.Decode(fImg2)

		//var width int = img1.Bounds().Dx()
		//var h int = img1.Bounds().Dy()

	var m = image.NewRGBA(image.Rect(0, 0,img1.Bounds().Dx(),img1.Bounds().Dy() ))
	mask := image.NewUniform(color.Alpha{150})
	draw.Draw(m, m.Bounds(), img1, image.Point{0,0}, draw.Src)


	draw.DrawMask(m, m.Bounds(), img3, image.Point{0,-400},mask,image.Point{-200,-100}, draw.Over)



	mask2 := image.NewUniform(color.Alpha{240})


	draw.DrawMask(m, m.Bounds(), img2, image.Point{-1*(img1.Bounds().Dx())+150,-450},mask2,image.Point{-200,-100}, draw.Over)

	//str := "clips/out"+strconv.Itoa(i)+".png"

	//toimg, _ := os.Create(str)
	toimg, _ := os.Create("news.png")
	defer toimg.Close()
	jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})
	//}()

}



