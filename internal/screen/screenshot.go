package screen

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"shotglass/internal/storage"
	"time"
)

func CreateScreenshot() ([]string, error) {
	n := screenshot.NumActiveDisplays()
	var fileNames []string
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		now := time.Now()
		fileName := fmt.Sprintf("screenshot_%d_%dx%d-%d.png", i, bounds.Dx(), bounds.Dy(), now.Unix())
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
	return fileNames, nil
}
