package main

import (
	"fmt"
	"os"
)

func Tree(path string, hasParent bool) {

	dirs, _ := os.ReadDir(path)
	for _, entry := range dirs {
		if entry.Name()[0] == '.' {
			continue
		}
		if hasParent {
			fmt.Printf("\t └── %v \n", entry.Name())
		} else {
			fmt.Printf("├──  %v\n", entry.Name())

		}
		Tree(path+entry.Name(), true)
	}

}
