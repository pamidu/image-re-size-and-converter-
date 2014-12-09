package main

import (
	"fmt"
	//"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	ImageFile := "01.jpg"
	fmt.Println("imag file :", ImageFile)
	dataofimage := ImageRead(ImageFile)
	fmt.Println("Reading image data ...")
	fmt.Println("Converting  image to png  ...")
	Formatpng(dataofimage)
	fmt.Println("Converting  image to gif  ...")
	Formatgif(dataofimage)

}

func ImageRead(ImageFile string) (image image.Image) {
	// open "test.jpg"
	file, err := os.Open(ImageFile)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return img
}

func Formatpng(img image.Image) {
	out, err := os.Create("converterdTOPNG.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" success...  ")

}

func Formatgif(img image.Image) {
	out, err := os.Create("converterdTOgif.gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var opt gif.Options
	opt.NumColors = 256
	err = gif.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" success... ")

}
