package main

import (
	"fmt"
	"os"
	path2 "path"
	"sort"
	"strings"
)

func Tree(path string, depth int) {
	node := "├──"
	lastNode := "└──"
	backBone := "│"

	dirContent, _ := os.ReadDir(path)

	lastidx := len(dirContent) - 1

	sort.Slice(dirContent, func(i, j int) bool {
		return dirContent[i].Name()[0] > dirContent[j].Name()[0]
	})
	for idx, content := range dirContent {
		if content.Name()[0] == '.' {
			continue
		}

		if idx != lastidx {
			if !content.IsDir() {
				fmt.Printf(" %v %v %v %v \n", backBone, strings.Repeat("\t", depth), node, content.Name())
			} else {
				fmt.Printf("%v %v %v \n", strings.Repeat("\t", depth), node, content.Name())

				Tree(path2.Join(path, content.Name()), depth+1)
			}
		} else {
			if !content.IsDir() {
				fmt.Printf(" %v %v %v %v \n", backBone, strings.Repeat("\t", depth), lastNode, content.Name())

			} else {
				fmt.Printf("%v %v %v \n", strings.Repeat("\t", depth), lastNode, content.Name())
				Tree(path2.Join(path, content.Name()), depth+1)
			}

		}
	}

}
