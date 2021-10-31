package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported originalFile")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds originalFile size")
	ErrUnknownOriginalFileSize = errors.New("original file size unknown")
	ErrorOpenFile   = errors.New("file open failed")
	ErrorCreateFile = errors.New("file create failed")
	originalFile    *os.File
	targetFile *os.File
)

//type fileChecker interface {
//	checkFile()
//}
//
//type originalFile struct {
//	originalFile    *os.File
//}
//
//func (f originalFile) check()  {
//	fileInfo, _ := f.Stat()
//	if ErrorOpenFile != nil {
//		if os.IsNotExist(ErrorOpenFile) {
//			return ErrorOpenFile
//		}
//	}
//}
//
//type targetFile struct {
//	targetFile *os.File
//}

func Copy(fromPath, toPath string, offset, limit int64) error {
	if filepath.Ext(fromPath) == "" {
		log.Printf("unsupported originalFile %v", fromPath)
		return ErrUnsupportedFile
	}
	originalFile, ErrorOpenFile = os.Open(fromPath)
	fileInfo, _ := originalFile.Stat()
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
