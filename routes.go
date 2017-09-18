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

func cdUp(w http.ResponseWriter, r *http.Request) {
	splitUp := strings.Split(Current, "/")
	Current = strings.Join(splitUp[:len(splitUp)-1], "/")

	dir := read()

	writeJSON(w, &dir)
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

func initial(w http.ResponseWriter, r *http.Request) {
	Current = Init

	dir := read()

	writeJSON(w, &dir)
}

func next(w http.ResponseWriter, r *http.Request) {
	afNext()
	writeJSON(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	thing := r.URL.Query()["item"][0]
	fullPath := Current + "/" + thing

	afKill()
	afPlay(fullPath)
	writeJSON(w, nil)
}

func playdir(w http.ResponseWriter, r *http.Request) {
	afKill()
	afPlay(Current)
	writeJSON(w, nil)
}

func readDir(w http.ResponseWriter, r *http.Request) {
	thing := r.URL.Query()["item"][0]
	fmt.Println(thing)

	Current = Current + "/" + thing

	dir := read()

	writeJSON(w, &dir)
}

func stop(w http.ResponseWriter, r *http.Request) {
	afKill()
	writeJSON(w, nil)
}
