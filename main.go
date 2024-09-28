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

	sources, err := core.LineSourceFactory{}.CreateInstance(findOptions.Path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, source := range sources {
		core.ProcessSource(source, findOptions)
	}
}
