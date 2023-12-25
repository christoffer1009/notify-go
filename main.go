package main

import (
	"io"
	"log"
	"notify/service"
	"os"
	"path/filepath"
)

func logsDirExists(logsDir string) bool {
	_, err := os.Stat(logsDir)
	return !os.IsNotExist(err)
}

func createLogsDirAndFile(logsDir, logsFile string) string {
	if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	logsPath := filepath.Join(logsDir, logsFile)

	_, err := os.Create(logsPath)
	if err != nil {
		log.Fatal(err)
	}

	return logsPath
}

func openLogs(logsPath string) *os.File {
	logs, err := os.OpenFile(logsPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return logs
}

func main() {
	url := "http://127.0.0.1:8080"
	logsDir := "logs"
	logsFile := "logs.log"
	logsPath := filepath.Join(logsDir, logsFile)

	if !logsDirExists(logsDir) {
		createLogsDirAndFile(logsDir, logsFile)
	}

	logs := openLogs(logsPath)
	defer logs.Close()

	multiWriter := io.MultiWriter(os.Stdout, logs)
	log.SetOutput(multiWriter)

	status, err := service.CheckStatus(url)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	log.Printf("Status: %s", status)

	if status == "online" {
		service.NotifyOnline()
	}
}
