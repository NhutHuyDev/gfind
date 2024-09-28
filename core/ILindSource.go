package core

import (
	"fmt"
	"strings"
)

type ILineSource interface {
	GetSourceName() string
	ReadLine() (Line, error)
	Open() error
	Close() error
}

func ProcessSource(source ILineSource, options FindOptions) {
	var isContain bool
	source = NewFilteredLineSource(source, func(l Line) bool {
		if !options.IsCaseSensitive {
			isContain = strings.Contains(strings.ToLower(l.text), strings.ToLower(options.StringToFind))
		} else {
			isContain = strings.Contains(l.text, options.StringToFind)
		}

		return (options.FindDontContain && !isContain) || (!options.FindDontContain && isContain)
	})

	err := source.Open()
	if err != nil {
		fmt.Println(err.Error())
	}

	if options.CountMode {
		CountMode(source)
	} else {
		PrintMode(source, options.ShowLineNumber)
	}

	err = source.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CountMode(source ILineSource) {
	count := 0
	for {
		_, err := source.ReadLine()
		if err != nil {
			break
		}

		count++
	}

	fmt.Printf("\n## %v: %v", source.GetSourceName(), count)
}

func PrintMode(source ILineSource, showLineNumber bool) {
	fmt.Printf("\n######### %v ######### \n", source.GetSourceName())
	for {
		line, err := source.ReadLine()
		if err != nil {
			break
		}

		PrintLine(line, showLineNumber)
	}
}
