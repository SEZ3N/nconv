package core

import (
	"github.com/SEZ3N/nconv/utils"
	"image"
	"image/color"
	"slices"
)

func GetLuminance(c color.Color) float64 {
	r, g, b, a := c.RGBA()
	return (0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)) / float64(a)
}

func Blend(blendFactor float64) color.Color {
	out := color.RGBA{
		R: uint8(float64(255) * blendFactor),
		G: uint8(float64(255) * blendFactor),
		B: uint8(float64(255) * blendFactor),
		A: uint8(255),
	}
	return out
}

func PickLum(picker []float64, bands []float64, lum float64) float64 {
	idx, _ := slices.BinarySearch(bands, lum)
	if idx == 0 {
		return picker[0]
	} else if idx >= len(bands)-1 {
		return picker[len(picker)-1]
	}
	return picker[idx-1]
}

func GenGreyscaleColorCoeffs(n uint) []float64 {
	out := make([]float64, 0)
	delta := 1.00 / float64(n-1)
	for i := 0; i < int(n); i++ {
		out = append(out, delta*float64(i))
	}
	return out
}

func InitLumArray(img image.Image) [][]float64 {
	var LumArray [][]float64

	for i := 0; i <= img.Bounds().Max.X; i++ {
		var temp []float64
		for j := 0; j <= img.Bounds().Max.Y; j++ {
			temp = append(temp, GetLuminance(img.At(i, j)))
		}
		LumArray = append(LumArray, temp)
	}
	return LumArray
}

func InitLuminanceBands(n uint, initalLum, finalLum float64) []float64 {
	bands := make([]float64, 0, n+1)
	offset := (finalLum - initalLum) / float64(n)
	for i := 0; i <= int(n); i++ {
		bands = append(bands, utils.Abs(initalLum+float64(i)*offset))
	}
	return bands
}
