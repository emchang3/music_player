package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func playFile(filePath string) {
	fmt.Printf("--- %v\n", filePath)

	cmd := exec.Command("afplay", filePath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func playDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fullPath := dir + "/" + file.Name()
		if file.IsDir() {
			playDir(fullPath)
		} else {
			if strings.Contains(fullPath, ".m4a") {
				playFile(fullPath)
			}
		}
	}
}

func main() {
	fmt.Println("PLAYING")

	fullPath := os.Args[1]

	if strings.Contains(fullPath, ".m4a") {
		playFile(fullPath)
	} else {
		playDir(fullPath)
	}
}
