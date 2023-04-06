package main

import (
	"os"
	"time"

	"github.com/MuhammadTalhaRao/zapfilerotation/constants"
)

type FileRotateWriter struct {
	Filename    string
	Interval    time.Duration
	file        *os.File
	lastRotated time.Time
	size        int
	MaxSize     int
}

/*
This package overwrite can be used as a pluging for zap logger to perform time and file based rotation
It instantiates file rotator which can rotate file based on specified time and size.
@params

	complete file path
	time interval duration
	max size of file

@returns

	file rotation instance

NOTE: One thing that needs to be kept in mind, files rotation criteria will only be checked when a log is written
*/
func NewTimeRotationWriter(filename string, interval time.Duration, maxSize int) *FileRotateWriter {
	return &FileRotateWriter{
		lastRotated: time.Now(),
		Filename:    filename,
		Interval:    interval,
		MaxSize:     maxSize * constants.MB,
	}
}

// rotates the file based on file size and time
func (w *FileRotateWriter) Write(output []byte) (int, error) {
	now := time.Now()

	// existing file size plus size of new log > max specified size
	sizeRotationRequired := w.size+len(output) > w.MaxSize

	// current time minus last rotation time > specified interval
	timeRotationRequired := now.Sub(w.lastRotated) > w.Interval

	// 1st condition is for time based rotation and 2nd is for size based rotation
	if timeRotationRequired || sizeRotationRequired {
		if err := w.rotate(); err != nil {
			return 0, err
		}

		// reset last rotation time
		if timeRotationRequired {
			w.lastRotated = time.Now()
		}
	}

	// this will create a new file in the begining when no file exists
	if w.file == nil {
		if err := w.rotate(); err != nil {
			return 0, err
		}
	}

	n, err := w.file.Write(output)
	w.size += n // keep track of file size

	return n, err
}

// renames old file and creates a new one
func (w *FileRotateWriter) rotate() error {
	if _, err := os.Stat(w.Filename); err == nil {
		backupFilename := w.Filename + "." + time.Now().Format("2006-01-02T15-04-05.000")
		if err := os.Rename(w.Filename, backupFilename); err != nil {
			return err
		}
	}

	file, err := os.Create(w.Filename)
	if err != nil {
		return err
	}

	w.file = file
	// reset file size
	w.size = 0

	return nil
}
