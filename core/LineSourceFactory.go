package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineSourceFactory struct {
}

func (LineSourceFactory LineSourceFactory) CreateInstance(path string) ([]ILineSource, error) {
	lineSource := make([]ILineSource, 0)

	var filePattern string
	var dirPath string

	if path == "" {
		return append(lineSource, NewConsoleLineSource()), nil
	} else {
		idx := strings.LastIndex(path, "/")

		if idx < 0 {
			dirPath = "."
			filePattern = path
		} else {
			dirPath = path[:idx]
			filePattern = path[idx+1:]
			if filePattern == "" {
				filePattern = "*"
			}
		}

		dir, err := os.Open(dirPath)
		if err != nil {
			return lineSource, fmt.Errorf("GFIND: '%v': No such file or directory 1", path)
		}
		defer dir.Close()

		files, err := dir.Readdir(-1)
		if err != nil {
			return lineSource, fmt.Errorf("GFIND: '%v': No such file or directory 2", path)
		}

		for _, file := range files {
			fileName := file.Name()

			matched, err := filepath.Match(filePattern, fileName)
			if err != nil {
				return lineSource, fmt.Errorf("GFIND: Invalid pattern: %v", filePattern)
			}

			if matched && !file.IsDir() {
				// fmt.Println(dirPath + "/" + file.Name())
				lineSource = append(lineSource, NewFileLineSource(dirPath+"/"+fileName, fileName))
			}
		}

		if len(lineSource) > 0 {
			return lineSource, nil
		} else {
			return lineSource, fmt.Errorf("GFIND: '%v': No such file or directory 3", path)
		}
	}
}
