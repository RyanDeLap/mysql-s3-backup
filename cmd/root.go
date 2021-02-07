package cmd

import (
	"fmt"
	"github.com/RyanDeLap/mysql-s3-backup/lib"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.000Z"

func Execute() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	fmt.Println("Starting backup of MySQL to S3...")

	fileName := time.Now().Format(timeLayout) + ".tar.gz"

	err = lib.DumpFiles(fileName)
	if err != nil {
		fmt.Println("Error generating file:", err)
	}

	fmt.Println("Successfully finished creating file for backup. Starting upload to S3.")

	location, err := lib.Upload(fileName)
	if err != nil {
		fmt.Println("Error uploading file:", err)
	}

	err = os.Remove(fileName)
	fmt.Println("Deleting temporary file...")
	if err != nil {
		fmt.Println("Error deleting temporary file:", err)
	}

	fmt.Println("Done! Uploaded file path: " + location)
}