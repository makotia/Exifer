package modules

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

func replace(s *tiff.Tag) string {
	return strings.ReplaceAll(s.String(), "\"", "")
}

func convertFNumber(s *tiff.Tag) float64 {
	numbers := strings.Split(replace(s), "/")
	var numerator, denominator int
	numerator, _ = strconv.Atoi(numbers[1])
	denominator, _ = strconv.Atoi(numbers[0])
	return float64(denominator) / float64(numerator)
}

// Decode : EXIFデータを返す
func Decode(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	camMaker, _ := x.Get(exif.Make)
	camModel, _ := x.Get(exif.Model)
	lensModel, _ := x.Get(exif.LensModel)
	exposureTime, _ := x.Get(exif.ExposureTime)
	iso, _ := x.Get(exif.ISOSpeedRatings)
	fNumber, _ := x.Get(exif.FNumber)
	fmt.Println("Camera:", replace(camMaker), replace(camModel))
	fmt.Println("Lens:", replace(lensModel))
	fmt.Println("ExposureTime:", replace(exposureTime))
	fmt.Println("ISO:", iso)
	fmt.Print("Aperture: ", "F", convertFNumber(fNumber))
}
