package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/NhutHuyDev/gfind/core"
)

func main() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChannel
		os.Exit(0)
	}()

	args := make([]string, 0)
	args = append(args, os.Args...)

	if len(args) == 1 {
		fmt.Println("GFIND: Parameter format not correct")
		return
	}

	findOptions, err := core.BuildOptions(args[1:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if findOptions.HelpMode {
		helpText := fmt.Sprintf(
			"Searches for a text string in a file or files.\n\n" +
				"gfind [-v] [-c] [-n] [-i] -f <\"string\"> [[drive:][path]filename[...]]\n\n" +
				"  -v         	Displays all lines NOT containing the specified string.\n" +
				"  -c         	Displays only the count of lines containing the string.\n" +
				"  -n         	Displays line numbers with the displayed lines.\n" +
				"  -i         	Ignores the case of characters when searching for the string.\n" +
				"  -f         	Required, defines the declaration for the string to search for\n" +
				"             	and the file or files to search in.\n" +
				"  <\"string\">   	Specifies the text string to find.\n" +
				"  [drive:][path]filename\n" +
				"             	Specifies a file or files to search.\n\n" +
				"If a path is not specified, FIND searches the text typed at the prompt\n" +
				"or piped from another command.",
		)
		fmt.Print(helpText)
		return
	}

	sources, err := core.LineSourceFactory{}.CreateInstance(findOptions.Path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, source := range sources {
		core.ProcessSource(source, findOptions)
	}
}
