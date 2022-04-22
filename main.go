package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"os"
	"shotglass/internal/conf"
	"shotglass/internal/screen"
	"shotglass/internal/storage"
	"shotglass/internal/util"
	"time"
)

func main() {
	cron := gocron.NewScheduler()
	err := cron.Every(1).Minutes().Do(screenShotAndUpload)
	if err != nil {
		fmt.Printf("Error occured on cron %s\n", err)
	}
	cron.Start()
	fmt.Println("JOB STARTED ...")
	screenShotAndUpload()
	select {}
}

func screenShotAndUpload() {
	success := false
	now := time.Now().Unix()
	files, err := util.WalkMatch(conf.AppConfig.Root, "*.txt")

	if err != nil {
		fmt.Printf("Error on getting root directory files: %s\n", err)
	}
	if len(files) > 0 {
		for _, file := range files {
			err = os.Rename(file, fmt.Sprintf("%s.%d", file, now))
			if err != nil {
				fmt.Printf("Error on renaming file: %s\n", err)
			}
		}
	}
	fileNames, err := screen.CreateScreenshot(now)
	if err != nil {
		fmt.Printf("Error on creating screenshots %s\n", err)
		return
	}
	fileNames = append(fileNames, files...)
	for _, fileName := range fileNames {
		uploadInfo, err := storage.PutObject(fileName)
		if err != nil {
			fmt.Printf("Error on file upload: %s\n", err)
			continue
		}
		success = true
		fmt.Printf("Screenshot uploaded successfully to %s\n", uploadInfo)
		err = os.Remove(fileName)
		if err != nil {
			fmt.Printf("Error occured on file remove: %s\n", err)
		}
	}
	if success {
		fmt.Println("Screenshot/Screenshots captured and uploaded successfully ...")
		return
	}
	fmt.Println("Screenshot failure")
}
