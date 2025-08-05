package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Tree(path string, hasParent bool) {

	dirs, _ := os.ReadDir(path)
	for _, entry := range dirs {
		if entry.Name()[0] == '.' {
			continue
		}
		if entry.IsDir() {
			fmt.Printf("\t └── %v \n", entry.Name())
			Tree(filepath.Join(path, entry.Name()), true)
		}

	}

}
