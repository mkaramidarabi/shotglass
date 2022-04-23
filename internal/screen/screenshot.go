package screen

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"shotglass/internal/storage"
)

func CreateScreenshot(now int64) ([]string, error) {
	n := screenshot.NumActiveDisplays()
	var fileNames []string
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d.png", now)
		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			return fileNames, err
		}
		err = png.Encode(file, img)
		if err != nil {
			return fileNames, err
		}
		_, err = storage.PutObject(fileName)
		if err != nil {
			return fileNames, err
		}
		fileNames = append(fileNames, fileName)
		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
	for _, file := range fileNames {
		fileName, err := Crop(file)
		if err != nil {
			fmt.Printf("Error occured croping file %s\n", fileName)
		}
	}
	return fileNames, nil
}
