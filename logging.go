package main

import (
	"io"
	"log"
	"os"
)

// custom log object
var (
	UploadLog *log.Logger
	ErrorLog  *log.Logger
)

func InitLogging(cfg Config) {
	// Inits config to logger objects
	if cfg.Flags.LogUpload {
		f, err := os.OpenFile(cfg.Paths.UploadLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			UploadLog = log.New(f, "UPLOAD: ", log.LstdFlags|log.Lshortfile) // only file
		}
	}

	f, err := os.OpenFile(cfg.Paths.ErrorLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		if cfg.Flags.LogError {
			writer := io.MultiWriter(os.Stdout, f) // console + file
			ErrorLog = log.New(writer, "ERROR: ", log.LstdFlags|log.Lshortfile)
		} else {
			ErrorLog = log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile) // only console
		}
	}
}
