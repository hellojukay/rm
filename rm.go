package main

import (
	"os"
)

func rm(file string) error {
	info, err := os.Stat(file)
	if err != nil {
		return err
	}
	if info.IsDir() {
		err = os.RemoveAll(file)
	} else {
		err = os.Remove(file)
	}
	if err != nil {
		return err
	}
	return nil
}
