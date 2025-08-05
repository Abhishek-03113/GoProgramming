package main

import "os"

func main() {
	homeDir, _ := os.UserHomeDir()
	path := homeDir + "/IdeaProjects/"
	Tree(path, "", false)
}
