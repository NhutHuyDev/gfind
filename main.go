package main

import (
	"fmt"
	"os"

	"github.com/NhutHuyDev/gfind/core"
)

func main() {
	args := make([]string, 0)
	args = append(args, os.Args...)

	if len(args) == 1 {
		fmt.Println("GFIND: Parameter format not correct")
		return
	}

	_, ok := core.BuildOptions(args[:1])
	if ok != nil {
		fmt.Println(ok.Error())
		return
	}
}
