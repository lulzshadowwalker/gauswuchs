package gauswuchs

import (
	"errors"
	"image"
	"image/color"
	"log"
	"math"
)

func Blur(m image.Image, kernelSize int, strength float64) image.Image {
	kernel, kernelSum, err := newKernel(kernelSize, int(strength))
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()
	out := image.NewRGBA(bounds)
	rad := kernelSize / 2

	for y := 0; y <= bounds.Dy(); y++ {
		for x := 0; x <= bounds.Dx(); x++ {
			var r, g, b, a float64

			for ky := range kernel {
				for kx := range kernel {
					rr, gg, bb, aa := decodeColor(m.At(x+kx-rad, y+ky-rad))
					r += float64(rr) * kernel[kx][ky]
					g += float64(gg) * kernel[kx][ky]
					b += float64(bb) * kernel[kx][ky]
					a += float64(aa) * kernel[kx][ky]
				}
			}

			r /= kernelSum
			g /= kernelSum
			b /= kernelSum
			a /= kernelSum

			out.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
		}
	}

	return out
}

func gauss(x, y, sd float64) float64 {
	return 1.0 / (2.0 * math.Pi * sd * sd) * math.Exp(-(x*x+y*y)/(2.0*sd*sd))
}

func decodeColor(c color.Color) (r, g, b, a uint8) {
	rr, gg, bb, aa := c.RGBA()
	return uint8(rr >> 8), uint8(gg >> 8), uint8(bb >> 8), uint8(aa >> 8)
}

func newKernel(size, strength int) (kernel [][]float64, sum float64, err error) {
	if size%2 == 0 {
		// so it can calculate the value of the pixel in the center
		return nil, -1, errors.New("kernel size has to be an odd number")
	}

	kernel = make([][]float64, size)
	for i := 0; i < size; i++ {
		kernel[i] = make([]float64, size)
	}

	rad := size / 2
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			kernel[x][y] = gauss(float64(x-rad), float64(y-rad), 100)
			sum += kernel[x][y]
		}
	}

	return
}
