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
	upperBound := mean + 2*stdDev
	lowerBound := mean - 2*stdDev

	// fallback values if values do not lie in [0,1] interval
	if lowerBound < 0 {
		mean = upperBound
		lowerBound = 0
		upperBound += 2 * stdDev
	}
	if upperBound > 1 {
		upperBound = 1
	}

	// partions the [lowerBound,upperBound] interval into n subintervals to be used for assigning the resultatnt greyscaleValues
	bands := core.InitLuminanceBands(n, lowerBound, upperBound)

	// Generate n equally spaced coefficients to be used to calulate the n Grey Scale Colors
	picker := core.GenGreyscaleColorCoeffs(n)

	for i := 0; i < len(lumArr); i++ {
		for j := 0; j < len(lumArr[0]); j++ {
			cl := core.PickLum(picker, bands, lumArr[i][j]) // PickLum performs binary search on bands to find in which interval the luminace lies
			out.Set(i, j, core.Blend(cl))
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
