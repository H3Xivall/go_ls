package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	entries, err := os.ReadDir(cwd)
	if err != nil {
		log.Fatal("Failed to read directory:", err)
	}

	var dirFiles []string
	for _, e := range entries {
		if !e.IsDir() {
			dirFiles = append(dirFiles, filepath.Join(cwd, e.Name()))
		}
	}

	if len(dirFiles) == 0 {
		fmt.Println("No Files found in this directory")
		return
	}
	sort.Strings(dirFiles)

	for i := 0; i < len(dirFiles); i++ {
		fmt.Printf(": %s :\n", dirFiles[i])
	}
}
