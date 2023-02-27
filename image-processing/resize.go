package main

import (
	"image"
	"image/color"
)

func resize(m image.Image, newBounds image.Rectangle) image.Image {
	res := image.NewRGBA64(image.Rect(newBounds.Min.X, newBounds.Min.Y, newBounds.Max.X, newBounds.Max.Y))
	for x := newBounds.Min.X; x <= newBounds.Max.X; x++ {
		for y := newBounds.Min.Y; y <= newBounds.Max.Y; y++ {
			c := m.At(x, y)
			r, g, b, a := c.RGBA()

			res.SetRGBA64(x, y, color.RGBA64{
				R: uint16(r),
				G: uint16(g),
				B: uint16(b),
				A: uint16(a),
			})
		}
	}

	return res
}
