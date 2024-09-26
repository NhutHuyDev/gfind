package main

import (
	"fmt"
	"os"
)

func main() {
	// Lấy các tham số dòng lệnh từ os.Args
	args := make([]string, 0)
	args = append(args, os.Args...)

	if len(args) == 1 {
		fmt.Println("GFIND: Parameter format not correct")
		return
	}

	_, ok := BuildOptions(args[:1])
	if ok != nil {
		fmt.Println(ok.Error())
		return
	}
}
