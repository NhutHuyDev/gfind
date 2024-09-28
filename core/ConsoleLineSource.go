package core

import (
	"bufio"
	"os"
	"strings"
)

type ConsoleLindSource struct {
	lineNumber int
	reader     *bufio.Reader
}

func (consoleFileSource *ConsoleLindSource) GetSourceName() string {
	return ""
}

func (consoleFileSource *ConsoleLindSource) ReadLine() (Line, error) {
	text, err := consoleFileSource.reader.ReadString('\n')
	if err != nil {
		return Line{}, err
	}

	consoleFileSource.lineNumber = consoleFileSource.lineNumber + 1

	return Line{LineNumber: consoleFileSource.lineNumber, text: strings.TrimSpace(text)}, nil
}

func (consoleFileSource *ConsoleLindSource) Open() error {
	return nil
}

func (consoleFileSource *ConsoleLindSource) Close() error {
	return nil
}

func NewConsoleLineSource() *ConsoleLindSource {
	return &ConsoleLindSource{lineNumber: 0,
		reader: bufio.NewReader(os.Stdin)}
}
