package main

import (
	"image"
	"image/color"
	"image/draw"
)

func main() {
	// TODO draw with svg from session creation
	img := image.NewRGBA(image.Rect(0, 0, 800, 600))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
}
