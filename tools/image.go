package tools

import (
	"bytes"
	"image"
	"image/draw"
	"log"
)

type IImage interface {
	Png2RGBA(png []byte) *image.RGBA
	Png2Image(png []byte) image.Image
}

type cimage struct{}

var Image IImage = &cimage{}

func (s *cimage) Png2RGBA(png []byte) *image.RGBA {
	img, _, _ := image.Decode(bytes.NewReader(png))
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	draw.Draw(rgbaImg, rgbaImg.Bounds(), img, bounds.Min, draw.Src)
	return rgbaImg
}

func (s *cimage) Png2Image(png []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(png))
	if err != nil {
		log.Fatal(err)
	}
	return img
}
