package tools

import (
	"bytes"
	"image"
	"image/draw"
)

type IImage interface {
	Png2RGBA(png []byte) (*image.RGBA, error)
	Png2Image(png []byte) (image.Image, error)
}

type cimage struct{}

var Image IImage = &cimage{}

func (s *cimage) Png2RGBA(png []byte) (*image.RGBA, error) {
	img, _, err := image.Decode(bytes.NewReader(png))
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	draw.Draw(rgbaImg, rgbaImg.Bounds(), img, bounds.Min, draw.Src)
	return rgbaImg, nil
}

func (s *cimage) Png2Image(png []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(png))
	return img, err
}
