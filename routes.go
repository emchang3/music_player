package main

import (
	"fmt"
	"html/template"
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

	dir := read()

	writeJSON(w, &dir)
}

func cdUp(w http.ResponseWriter, r *http.Request) {
	splitUp := strings.Split(Current, "/")
	Current = strings.Join(splitUp[:len(splitUp)-1], "/")

	dir := read()

	writeJSON(w, &dir)
}

func initial(w http.ResponseWriter, r *http.Request) {
	Current = Init

	dir := read()

	writeJSON(w, &dir)
}

func play(w http.ResponseWriter, r *http.Request) {
	thing := r.URL.Query()["item"][0]
	fullPath := Current + "/" + thing

	if strings.Contains(fullPath, ".m4a") {
		playFile(fullPath)
	} else {
		playDir(fullPath)
	}

	writeJSON(w, nil)
}

func pause(w http.ResponseWriter, r *http.Request) {
	afStop()
	writeJSON(w, nil)
}

func cont(w http.ResponseWriter, r *http.Request) {
	afCont()
	writeJSON(w, nil)
}
