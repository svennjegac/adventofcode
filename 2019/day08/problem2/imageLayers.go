package main

import (
	"image"
	"image/png"
	"os"

	"github.com/svennjegac/adventofcode/2019/day08"
)

const (
	width  = 25
	height = 6
)

func main() {
	img, err := day08.Image()
	if err != nil {
		panic(err)
	}

	rgbImage := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < width*height; i++ {
		for j := i; ; j += width * height {
			if img[j] != 2 {
				for k := i * 4; k < i*4+4; k++ {
					rgbImage.Pix[k] = uint8(img[j] * 255)
				}
				break
			}
		}
	}

	output, err := os.Create("2019/day08/problem2/image.png")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	png.Encode(output, rgbImage)
}
