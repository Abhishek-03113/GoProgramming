package main

import (
	"fmt"
	"os"
	path2 "path"
)

func Tree(path string, prefix string, isLast bool) {
	dirContent, err := os.ReadDir(path)
	if err != nil {
		return
	}

	visible := []os.DirEntry{}
	for _, val := range dirContent {
		if val.Name()[0] != '.' {
			visible = append(visible, val)
		}
	}

	for idx, content := range visible {
		isLastItem := idx == len(visible)-1

		if isLastItem {
			fmt.Printf("%s└── %s\n", prefix, content.Name())
		} else {
			fmt.Printf("%s├── %s\n", prefix, content.Name())
		}

		if content.IsDir() {
			var newPrefix string
			if isLastItem {
				newPrefix = prefix + "    "
			} else {
				newPrefix = prefix + "│   "
			}
			Tree(path2.Join(path, content.Name()), newPrefix, isLastItem)
		}
	}
}
