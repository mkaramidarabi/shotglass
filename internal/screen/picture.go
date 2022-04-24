package screen

import (
	"fmt"
	"github.com/oliamb/cutter"
	"image"
	"image/png"
	"os"
)

func Crop(filename string) (string, error) {
	f, err := os.Open(filename)
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {

		}
	}(f)
	if err != nil {
		fmt.Printf("Error occured reading file %s\n", err)
		return "", err
	}
	img, _, err := image.Decode(f)
	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  450,                 // height in pixel or Y ratio(see Ratio Option below)
		Width:   1110,                // width in pixel or X ratio
		Mode:    cutter.TopLeft,      // Accepted Mode: TopLeft, Centered
		Anchor:  image.Point{0, 110}, // Position of the top left point
		Options: 0,                   // Accepted Option: Ratio
	})

	f2, err := os.Create(filename)
	defer func(f2 *os.File) {
		err = f2.Close()
		if err != nil {

		}
	}(f2)
	if err != nil {
		fmt.Printf("Error occured creating file %s\n", err)
		return "", err
	}
	err = png.Encode(f2, cImg)
	if err != nil {
		fmt.Printf("Error occured encoding file %s\n", err)
		return "", err
	}
	return filename, nil
}
