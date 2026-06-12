package core
import (
	"image"
	"image/color"
	"slices"

	"github.com/SEZ3N/nconv/utils"
)

func getLuminance(c color.Color) float64 {
	r, g, b, a := c.RGBA()
	return (0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)) / float64(a)
}

func LumFromCoeff(blendFactor float64) color.RGBA {
	col := uint8(255 * blendFactor)
	return color.RGBA{
		R: col,
		G: col,
		B: col,
		A: 255,
	}
}

func PickLumCoeff(picker []float64, bands []float64, lum float64) float64 {
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

func InitLumArray(img image.Image) []float64 {
	switch img.(type) {
	case *image.YCbCr:
		return initLumArrYCbCr(img)
	default:
		return initLumArrDefault(img)
	}
}

func InitLuminanceBands(n uint, initalLum, finalLum float64) []float64 {
	bands := make([]float64, 0, n+1)
	offset := (finalLum - initalLum) / float64(n)
	for i := 0; i <= int(n); i++ {
		bands = append(bands, utils.Abs(initalLum+float64(i)*offset))
	}
	return bands
}

func initLumArrYCbCr(img image.Image) []float64 {
	rect := img.Bounds()
	w,h := img.Bounds().Dx(),img.Bounds().Dy()
	lumArray := make([]float64,w*h)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			idx := img.(*image.YCbCr).YOffset(x, y)
			luma := img.(*image.YCbCr).Y[idx]
			lumArrIdx := (y - rect.Min.Y) * w + x - rect.Min.X
			lumArray[lumArrIdx] = float64(luma)/255.00
		}
	}
	return lumArray
}

func initLumArrDefault(img image.Image) []float64 {
	rect := img.Bounds()
	w,h := img.Bounds().Dx(),img.Bounds().Dy()
	lumArray := make([]float64,w*h)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			luma := getLuminance(img.At(x,y))
			lumArrIdx := (y - rect.Min.Y) * w + x - rect.Min.X
			lumArray[lumArrIdx] = float64(luma)/255.00
		}
	}
	return lumArray
}

