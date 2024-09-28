package core

import (
	"fmt"
	"strings"
)

type FindOptions struct {
	FindDontContain bool   // "/v"
	CountMode       bool   // "/c"
	ShowLineNumber  bool   // "/n"
	IsCaseSensitive bool   // "/i"
	StringToFind    string // "<"string">""
	Path            string // "[<drive>:][<path>]<filename>"
	HelpMode        bool   // "/?"
}

func NewFileOptions() FindOptions {
	return FindOptions{
		FindDontContain: false,
		CountMode:       false,
		ShowLineNumber:  false,
		IsCaseSensitive: true,
		StringToFind:    "",
		Path:            "",
		HelpMode:        false,
	}
}

func BuildOptions(optionArgs []string) (FindOptions, error) {
	options := NewFileOptions()

	isOption := true
	for _, arg := range optionArgs {
		if isOption {
			if arg == "-v" {
				options.FindDontContain = true
			} else if arg == "-c" {
				options.CountMode = true
			} else if arg == "-n" {
				options.ShowLineNumber = true
			} else if arg == "-i" {
				options.IsCaseSensitive = false
			} else if arg == "-help" {
				options.HelpMode = true
			} else if arg == "-f" {
				isOption = false
				continue
			} else {
				return options, fmt.Errorf("GFIND: Parameter format not correct")
			}
		} else {
			if options.StringToFind == "" {
				options.StringToFind = arg
			} else if options.Path == "" {
				path := strings.ReplaceAll(arg, "\\", "/")
				options.Path = path
			}
		}
	}

	return options, nil
}
