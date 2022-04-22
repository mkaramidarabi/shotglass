package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
	"shotglass/internal/screen"
	"shotglass/internal/storage"
)

func main() {
	cron := cron.New()
	cron.AddFunc("@every 1m", screenShotAndUpload)
	cron.Start()
	fmt.Println("JOB STARTED ...")
	screenShotAndUpload()
	select {}
}

func screenShotAndUpload() {
	success := false
	fileNames, err := screen.CreateScreenshot()
	if err != nil {
		fmt.Printf("Error on creating screenshots %s\n", err)
		return
	}
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
