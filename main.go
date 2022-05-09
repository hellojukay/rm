package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	for _, file := range os.Args[1:] {
		if err := rm(file); err != nil {
			os.Exit(1)
			fmt.Printf("remove %s error %s", file, err.Error())
		}
	}
}
