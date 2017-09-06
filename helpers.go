package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func read() []Item {
	var dir []Item

	files, err := ioutil.ReadDir(Current)
	if err != nil {
		fmt.Println(err)
		return dir
	}

	for _, file := range files {
		fileType := ""
		if file.IsDir() {
			fileType = "dir"
		} else {
			if strings.Contains(file.Name(), ".m4a") {
				fileType = "file"
			}
		}

		if fileType == "dir" || fileType == "file" {
			newFile := Item{Name: file.Name(), Type: fileType}
			dir = append(dir, newFile)
		}
	}

	return dir
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

func playFile(filePath string) {
	cmd := exec.Command("afplay", filePath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
