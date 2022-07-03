package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	width := 300
	height := 300

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{60, 150, 200, 0xff}
	another := color.RGBA{200, 80, 100, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, cyan)
			//switch {
			//case x < width/2 && y < height/2: // upper left quadrant
			//	img.Set(x, y, cyan)
			//case x >= width/2 && y >= height/2: // lower right quadrant

			if x < y {
				img.Set(x, y, color.Black)
			}
			if y > 0 && x*x == y {
				img.Set(x, y, color.White)
			}

			if math.Sin(float64(x)) > math.Exp(float64(y+1)) {
				img.Set(x, y, color.White)
			}
			if x*x+y*y < 7690 {
				img.Set(x, y, color.White)
			}
			if (width-x)*(width-x)+(height-y)*(height-y) < 27690 {
				img.Set(x, y, another)
			}
			// Use zero value.

		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
