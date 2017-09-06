package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"strings"
)

// Item ... Directory item.
type Item struct {
	Name string
	Type string
}

type Nil struct{}

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

func play(w http.ResponseWriter, r *http.Request) {
	thing := r.URL.Query()["item"][0]
	fmt.Println(thing)
	fullPath := Current + "/" + thing
	fmt.Println(fullPath)

	if strings.Contains(fullPath, ".m4a") {
		playFile(fullPath)
	} else {
		playDir(fullPath)
	}

	js, err := json.Marshal(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	return
}

func pause(w http.ResponseWriter, r *http.Request) {
	// killall -next
	cmd := exec.Command("killall", "-next", "afplay")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	return
}

func cont(w http.ResponseWriter, r *http.Request) {
	// killall -next
	cmd := exec.Command("killall", "-CONT", "afplay")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	return
}

func next(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("killall", "afplay")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	return
}
