package nconv

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/SEZ3N/nconv/core"
	"github.com/SEZ3N/nconv/utils"
)

func Convert(img image.Image, n uint) image.Image {
	out := image.NewRGBA(image.Rect(img.Bounds().Min.X, img.Bounds().Min.Y, img.Bounds().Max.X, img.Bounds().Max.Y))
	lumArr := core.InitLumArray(img)
	
	mean := utils.MeanLuminosity(lumArr) // luminosity mean
	stdDev := utils.StdDev(lumArr, mean) // luminosity stddev

	// the lines below calculates the range where 95% of the luminosity values in the image lie
	// this gives better result when most of the luminosity values in the image are either on the darker or ligher side
	lowerBound,upperBound := utils.CalcBounds(mean,stdDev)
	// partions the [lowerBound,upperBound] interval into n subintervals to be used for assigning the resultatnt greyscaleValues
	bands := core.InitLuminanceBands(n, lowerBound, upperBound)

	// Generate n equally spaced coefficients to be used to calulate the n Grey Scale Colors
	picker := core.GenGreyscaleColorCoeffs(n)

	rect := img.Bounds()
	w := img.Bounds().Dx()

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			idx := (y - rect.Min.Y) * w + (x - rect.Min.X)
			coeff := core.PickLumCoeff(picker, bands, lumArr[idx]) // PickLum performs binary search on bands to find in which interval the luminace lies
			out.SetRGBA(x, y, core.LumFromCoeff(coeff))
		}	
	}
	return out
}

func ConvertAndWrite(img image.Image, n uint, writer io.Writer, quality int) error {
	out := Convert(img, n)
	e := jpeg.Encode(writer, out, &jpeg.Options{Quality: quality})
	if e != nil {
		return e
	}
	return nil
}


