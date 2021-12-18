package main

import (
	"errors"
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
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var (
		originalFile *os.File
		targetFile   *os.File
		bufLen       int64
	)

	if filepath.Ext(fromPath) == "" || filepath.Ext(fromPath) != filepath.Ext(toPath) {
		log.Printf("unsupported originalFile %v", fromPath)
		return ErrUnsupportedFile
	}
	if offset < 0 {
		return ErrorNegativeOffset
	}

	originalFile, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(ErrorOpenFile) {
			return ErrorOpenFile
		}
	}

	fileInfo, err := originalFile.Stat()
	if err != nil {
		if os.IsNotExist(ErrorOpenFile) {
			return ErrorOpenFile
		}
	}
	defer originalFile.Close()

	if fileInfo.Size() <= limit {
		limit = fileInfo.Size()
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

	defer targetFile.Close()

	bar := pb.Full.Start64(limit)

	var buf []byte

	switch limit {
	case 0:
		read, err := io.ReadAll(originalFile)
		buf = read[offset:]

		if err != nil {
			log.Printf("failed to read original file: %v", err)
			return err
		}
	default:
		if offset+limit > fileInfo.Size() {
			bufLen = fileInfo.Size() - offset
		} else {
			bufLen = limit
		}
		buf = make([]byte, bufLen)

		originalFile.ReadAt(buf, offset)

		bar.Finish()
	}

	_, err = targetFile.Write(buf)
	if err != nil {
		log.Printf("failed to write: %v", err)
		return err
	}

	return nil
}
