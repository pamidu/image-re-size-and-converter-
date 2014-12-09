package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	ImageFile := "01.jpg"
	fmt.Println("imag file :", ImageFile)
	dataofimage := ImageRead(ImageFile)
	fmt.Println("Reading image data ...")
	fmt.Println("Resizing image ...")
	Resize(dataofimage)

}

func ImageRead(ImageFile string) (image image.Image) {
	// open "test.jpg"
	file, err := os.Open("01.jpg")
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

func Resize(img image.Image) {

	//https://github.com/nfnt/resize
	m := resize.Resize(5000, 5000, img, resize.NearestNeighbor)

	out, err := os.Create("test_qsized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	errors := jpeg.Encode(out, m, nil)
	if errors != nil {
		fmt.Println("somet thing goin wrong ")
		os.Exit(0)
	}
	fmt.Println("success...")

}
