package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	file := os.Args[1]
	info, err := os.Stat(file)
	if err != nil {
		os.Exit(1)
	}
	if info.IsDir() {
		err = os.RemoveAll(file)
	} else {
		err = os.Remove(file)
	}
	if err != nil {
		fmt.Printf("remove %s error %s", file, err.Error())
		os.Exit(1)
	}
}
