package main

import (
	"errors"
	"io"
	"log"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported originalFile")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds originalFile size")
	originalFile os.File
	targetFile os.File

)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.
	originalFile, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Not found
		}
	}

	targetFile, err := os.Create(toPath)

	buf := make([]byte, limit)

	// Check offset and limit
	for offset < limit {
		read, err := originalFile.Read(buf[offset:])

		offset += int64(read)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("failed to read original file: %v", err)
		}


	}



	defer originalFile.Close()
	defer targetFile.Close()


	return nil
}
