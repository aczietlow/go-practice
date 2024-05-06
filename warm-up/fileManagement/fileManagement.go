package main

import (
	"fmt"
	"os"
)

func main() {
	// Hard code file path for now.
	path := "/home/aczietlow/Projects/go-practice/warm-up/fileManagement"
	// @TODO stat path for valid filepath
	// @TODO stat path for directory or file
	list(path)
}

func list(path string) {
	dirs, err := os.ReadDir(path)

	if err != nil {
		fmt.Println(err)
	}

	// Range over directory entries.
	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}
}
