package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileLineSource struct {
	path       string
	fileName   string
	file       *os.File
	reader     *bufio.Reader
	lineNumber int
}

func (fileLineSource *FileLineSource) GetSourceName() string {
	return fileLineSource.fileName
}

func (fileLineSource *FileLineSource) ReadLine() (Line, error) {
	if fileLineSource.reader == nil {
		return Line{}, fmt.Errorf("GFIND: Cannot open file %v", fileLineSource.path)
	}

	text, err := fileLineSource.reader.ReadString('\n')
	if err != nil {
		return Line{}, err
	}

	fileLineSource.lineNumber = fileLineSource.lineNumber + 1

	return Line{LineNumber: fileLineSource.lineNumber, text: strings.TrimSpace(text)}, nil
}

func (fileLineSource *FileLineSource) Open() error {
	file, err := os.Open(fileLineSource.path)
	if err != nil {
		return err
	}
	fileLineSource.file = file
	fileLineSource.reader = bufio.NewReader(fileLineSource.file)
	fileLineSource.lineNumber = 0
	return nil
}

func (fileLineSource *FileLineSource) Close() error {
	if fileLineSource.file != nil {
		err := fileLineSource.file.Close()
		if err != nil {
			return err
		}
		fileLineSource.file = nil
	}
	return nil
}

func NewFileLineSource(path string, fileName string) *FileLineSource {
	return &FileLineSource{path: path, fileName: fileName}
}
