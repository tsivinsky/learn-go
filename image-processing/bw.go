package main

import (
	"image"
	"image/color"
)

func toBW(m image.Image) image.Image {
	bounds := m.Bounds()

	bw := image.NewRGBA64(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for x := 0; x <= bounds.Max.X; x++ {
		for y := 0; y <= bounds.Max.Y; y++ {
			c := m.At(x, y)
			r, g, b, a := c.RGBA()

			gray := uint16(rgbToBw(r, g, b))

			bw.SetRGBA64(x, y, color.RGBA64{
				R: gray,
				G: gray,
				B: gray,
				A: uint16(a),
			})
		}
	}

	return bw
}
