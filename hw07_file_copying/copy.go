package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile         = errors.New("unsupported originalFile")
	ErrOffsetExceedsFileSize   = errors.New("offset exceeds originalFile size")
	ErrUnknownOriginalFileSize = errors.New("original file size unknown")
	ErrorOpenFile              = errors.New("file open failed")
	ErrorCreateFile            = errors.New("file create failed")
	ErrorNegativeOffset        = errors.New("offset can not be negative")
	originalFile               *os.File
	targetFile                 *os.File
	bufLen                     int64
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if filepath.Ext(fromPath) == "" || filepath.Ext(fromPath) != filepath.Ext(toPath) {
		log.Printf("unsupported originalFile %v", fromPath)
		return ErrUnsupportedFile
	}
	if offset < 0 {
		return ErrorNegativeOffset
	}

	originalFile, ErrorOpenFile = os.Open(fromPath)
	fileInfo, _ := originalFile.Stat()

	if fileInfo.Size() <= limit {
		limit = fileInfo.Size()
	}

	if ErrorOpenFile != nil {
		if os.IsNotExist(ErrorOpenFile) {
			return ErrorOpenFile
		}
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	if fileInfo.Size() == 0 {
		return ErrUnknownOriginalFileSize
	}

	targetFile, ErrorCreateFile = os.Create(toPath)

	if ErrorCreateFile != nil {
		if os.IsNotExist(ErrorCreateFile) {
			return ErrorCreateFile
		}
	}
	defer originalFile.Close()
	defer targetFile.Close()

	bar := pb.Full.Start64(limit)

	switch limit {
	case 0:
		read, err := io.ReadAll(originalFile)
		copied := read[offset:]

		if err != nil {
			fmt.Println(err)
			log.Panicf("failed to read original file: %v", err)
		}

		_, err = targetFile.Write(copied)
		if err != nil {
			log.Panicf("failed to write: %v", err)
		}
	default:
		bufLen = fileInfo.Size()
		if offset+limit > fileInfo.Size() {
			bufLen = fileInfo.Size() - offset
		} else {
			bufLen = limit
		}
		buf := make([]byte, bufLen)

		originalFile.ReadAt(buf, offset)

		_, err := targetFile.Write(buf)
		if err != nil {
			log.Panicf("failed to write: %v", err)
		}
		bar.Finish()
	}

	return nil
}
