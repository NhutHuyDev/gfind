package core

import "fmt"

type FindOptions struct {
	FindDontContain  bool   // "/v"
	CountMode        bool   // "/c"
	ShowLineNumber   bool   // "/n"
	IsCaseSensitive  bool   // "/i"
	SkipOfflineFiles bool   // "[/off[line]]"
	StringToFind     string // "<"string">""
	Path             string // "[<drive>:][<path>]<filename>"
	HelpMode         bool   // "/?"
}

func NewFileOptions() FindOptions {
	return FindOptions{
		FindDontContain:  false,
		CountMode:        false,
		ShowLineNumber:   false,
		IsCaseSensitive:  true,
		SkipOfflineFiles: true,
		StringToFind:     "",
		Path:             "",
		HelpMode:         false,
	}
}

func BuildOptions(optionArgs []string) (FindOptions, error) {
	options := NewFileOptions()

	isFlag := true
	for _, arg := range optionArgs {
		if isFlag {
			if arg == "-v" {
				options.FindDontContain = true
			} else if arg == "-c" {
				options.CountMode = true
			} else if arg == "-n" {
				options.ShowLineNumber = true
			} else if arg == "-i" {
				options.IsCaseSensitive = false
			} else if arg == "-off" || arg == "-offline" {
				options.SkipOfflineFiles = false
			} else if arg == "-help" {
				options.HelpMode = true
			} else if arg == "--" {
				isFlag = false
				continue
			} else {
				return options, fmt.Errorf("GFIND: Parameter format not correct")
			}
		} else {
			if options.StringToFind == "" {
				options.StringToFind = arg
			} else if options.Path == "" {
				options.Path = arg
			}
		}
	}

	return options, nil
}
