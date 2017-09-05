package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

// Item ... Directory item.
type Item struct {
	Name string
	Type string
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	t, err := template.ParseFiles("views/index.gohtml")
	if err != nil {
		return
	}

	t.Execute(w, nil)
	return
}

func readDir(w http.ResponseWriter, r *http.Request) {
	thing := r.URL.Query()["item"][0]
	fmt.Println(thing)

	Current = Current + "/" + thing
	fmt.Println(Current)

	dir := read()

	js, err := json.Marshal(dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func cdUp(w http.ResponseWriter, r *http.Request) {
	splitUp := strings.Split(Current, "/")
	Current = strings.Join(splitUp[:len(splitUp)-1], "/")
	fmt.Println(Current)

	dir := read()

	js, err := json.Marshal(dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func initial(w http.ResponseWriter, r *http.Request) {
	Current = Init

	dir := read()

	js, err := json.Marshal(dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
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
