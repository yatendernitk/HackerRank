package main

import (
	"github.com/fogleman/gg"
	"log"
	"strconv"
	"fmt"
	"time"
	//"runtime"
	//"sync"
	//"sync"
	"runtime"
)
//var wg sync.WaitGroup
var l = 0
func main() {
	runtime.GOMAXPROCS(4)
	start := time.Now()
	for i:=0; i< 500 ; i++ {
		fmt.Println(i)
		drawImg(i)
	}
	//wg.Wait()
	elapsed := time.Since(start)
	log.Println(l)
	log.Printf("Binomial took %s", elapsed)
	nm := ""
	fmt.Scanf("%s",&nm)

}

func drawImg(i int) {
	l++
	//wg.Add(1)
	go func(){
		const S = 1024
		im, err := gg.LoadImage("src.jpg")
		if err != nil {
			log.Fatal(err)
		}

		dc := gg.NewContext(S, S)
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		dc.SetRGB(0, 0, 0)
		if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 96); err != nil {
			panic(err)
		}
		dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
		dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
		dc.DrawImage(im, 0, 0)
		dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
		dc.Clip()
		//dc.DrawImage(im, 0, 0)
		str := "clips/out"+strconv.Itoa(i)+".png"
		dc.SavePNG(str)
		//wg.Done()
	}()

}