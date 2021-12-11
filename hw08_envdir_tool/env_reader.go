package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	environment := make(Environment, len(files))

	if err != nil {
		log.Print(err)
		return nil, err
	}

	for _, file := range files {
		envVal := new(EnvValue)
		file, err := os.Open(dir + "/" + file.Name())
		if err != nil {
			log.Print(err)
			return nil, err
		}

		fi, err := file.Stat()
		if err != nil {
			log.Print(err)
			return nil, err
		}

		if fi.Size() == 0 {
			envVal.Value = ""
			envVal.NeedRemove = true
			environment[filepath.Base(file.Name())] = *envVal
			break
		}

		reader := bufio.NewReader(file)

		line, _, err := reader.ReadLine()
		if err != nil {
			log.Print(err)
			return nil, err
		}

		cleanStr := bytes.ReplaceAll(line, []byte{0}, []byte("\n"))
		cleanStr = []byte(strings.TrimRight(string(cleanStr), " \t"))

		envVal.Value = string(cleanStr)

		environment[filepath.Base(file.Name())] = *envVal
	}

	return environment, nil
}
