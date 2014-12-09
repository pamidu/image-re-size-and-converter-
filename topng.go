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
	ImageFilejpg := "01.jpg"
	ImageFilegif := "01.gif"
	fmt.Println("jpg imag file :", ImageFilejpg, "\n gif imag file :", ImageFilegif)
	//read image file
	fmt.Println("Reading image data ...")
	dataofimagejpg := ImageRead(ImageFilejpg, "jpeg")
	dataofimagegif := ImageRead(ImageFilegif, "gif")

	fmt.Println("Converting  image jpeg to png  ...")
	Formatpng(dataofimagejpg, "jpegTOpng.png")
	fmt.Println("Converting  image gif to png  ...")
	Formatpng(dataofimagegif, "gifTOpng")
	//fmt.Println("Converting  image to gif  ...")
	//Formatgif(dataofimage)

}

func ImageRead(ImageFile string, format string) (image image.Image) {
	if format == "jpeg" {
		file, err := os.Open(ImageFile)
		if err != nil {
			log.Fatal(err)
		}
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		return img

	} else if format == "gif" {
		file, err := os.Open(ImageFile)
		if err != nil {
			log.Fatal(err)
		}
		img, err := gif.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		return img
	}
	return nil
}

func Formatpng(img image.Image, name string) {
	out, err := os.Create(name)
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
