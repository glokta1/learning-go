package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	img [][]uint8
	w   int
	h   int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{0, img.w}, image.Point{0, img.h}}
}

func (img Image) At(x, y int) color.Color {
	v := img.img[x][y]
	return color.RGBA{v, v, 255, 255}
}

func Pic(dx, dy int) Image {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		// initialize each row of pic with []uint8 slice of dx size
		pic[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			pic[i][j] = uint8((i + j) / 2)
		}
	}

	return Image{pic, dx, dy}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
