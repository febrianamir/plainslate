package lib

import (
	"log"
	"os"
	"path/filepath"
)

// Log file
var logFile *os.File

func InitLogger() {
	var err error

	logPath := filepath.Join("debug", "debug.log")
	logFile, err = os.OpenFile(
		logPath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)
}

func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}
