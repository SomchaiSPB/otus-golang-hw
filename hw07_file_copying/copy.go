package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported originalFile")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds originalFile size")
	ErrorOpenFile = errors.New("file open failed")
	ErrorreateFile = errors.New("file create failed")
	originalFile *os.File
	targetFile *os.File
	tmpFile *os.File
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	originalFile, ErrorOpenFile = os.Open(fromPath)
	if ErrorOpenFile != nil {
		if os.IsNotExist(ErrorOpenFile) {
			return ErrorOpenFile
		}
	}
	targetFile, ErrorreateFile = os.Create(toPath)
	if ErrorreateFile != nil {
		if os.IsNotExist(ErrorreateFile) {
			return ErrorreateFile
		}
	}
	defer originalFile.Close()
	defer targetFile.Close()

	// Create tmp filename from original fileinfo

	//fileInfo := os.FileInfo()
	//
	//// TODO finish this
	//if offset > os.FileInfo().Size() {
	//	return ErrOffsetExceedsFileSize
	//}

	if limit > 0 {
		_, err := io.CopyN(targetFile, originalFile, limit)
		if err != nil {
			log.Panicf("failed to read original file: %v", err)
		}
	}


	if offset == 0 && limit == 0 {

		//read, err := io.ReadAll(originalFile)
		originalFile, err := ioutil.ReadFile(fromPath)

		if err != nil {
			log.Panicf("failed to read original file: %v", err)
		}

		//written, err := targetFile.Write(originalFile)
		err = ioutil.WriteFile(toPath, originalFile, 0644)
		if err!= nil {
			log.Panicf("failed to write: %v", err)
		}

	} else {
		buf := make([]byte, limit)
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

		_, err := targetFile.Write(buf)
		if err!= nil {
			log.Panicf("failed to write: %v", err)
		}

	}

	return nil
}
