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
		fileName, err := createScreenshotByDisplayNumber(i, now)
		if err != nil {
			continue
		}
		fileNames = append(fileNames, fileName)
	}
	for _, file := range fileNames {
		fileName, err := Crop(file)
		if err != nil {
			fmt.Printf("Error occured croping file %s\n", fileName)
		}
	}
	return fileNames, nil
}

func createScreenshotByDisplayNumber(i int, now int64) (string, error) {
	bounds := screenshot.GetDisplayBounds(i)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("p_%d.png", now)
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return "", err
	}
	err = png.Encode(file, img)
	if err != nil {
		return "", err
	}
	_, err = storage.PutObject(fileName)
	if err != nil {
		return "", err
	}
	fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
