package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("images/brasil.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	processedImage := processImage(img)

	outputFile, err := os.Create("output/image_with_inverted_colors.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, processedImage, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func processImage(img image.Image) image.Image {
	bounds := img.Bounds()
	processedImage := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			originalColor := img.At(x, y)

			r, g, b, a := originalColor.RGBA()
			newColor := color.RGBA{
				R: 255 - uint8(r>>8),
				G: 255 - uint8(g>>8),
				B: 255 - uint8(b>>8),
				A: uint8(a >> 8),
			}

			processedImage.Set(x, y, newColor)
		}
	}

	return processedImage
}
