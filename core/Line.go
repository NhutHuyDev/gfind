package core

import "fmt"

type Line struct {
	LineNumber int
	text       string
}

func PrintLine(line Line, showLineNumber bool) {
	if showLineNumber {
		fmt.Printf("[%v] %v \n", line.LineNumber, line.text)
	} else {
		fmt.Printf("%v \n", line.text)
	}
}
