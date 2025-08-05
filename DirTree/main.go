package main

import "os"

func main() {
	homeDir, _ := os.UserHomeDir()
	path := homeDir + "/GoProgramming/"
	Tree(path, 0)
}
