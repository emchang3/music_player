package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

func afKill() {
	cmd := exec.Command("shell_commands/music_control.sh", "1")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func afNext() {
	cmd := exec.Command("shell_commands/music_control.sh", "4")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func afPlay(path string) {
	cmd := exec.Command("shell_commands/play_music", path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(out.String())
}

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

func writeJSON(w http.ResponseWriter, item *[]Item) {
	js, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
