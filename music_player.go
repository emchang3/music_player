package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/NYTimes/gziphandler"
	"github.com/joho/godotenv"
)

// Current directory.
var Current = "/Users/emchang3/Music"

// Init directory.
var Init = "/Users/emchang3/Music"

func routeHandler() {
	fs := http.FileServer(http.Dir("public"))
	nopref := http.StripPrefix("/public/", fs)
	filesGz := gziphandler.GzipHandler(nopref)
	http.Handle("/public/", filesGz)

	// activatorGz := gziphandler.GzipHandler(http.HandlerFunc(fs2))
	// http.Handle("/432FB6766878ED13CC007C095B54B76A.txt", activatorGz)

	cdUpGz := gziphandler.GzipHandler(http.HandlerFunc(cdUp))
	// indexGz := gziphandler.GzipHandler(http.HandlerFunc(index))
	initialGz := gziphandler.GzipHandler(http.HandlerFunc(initial))
	nextGz := gziphandler.GzipHandler(http.HandlerFunc(next))
	playGz := gziphandler.GzipHandler(http.HandlerFunc(play))
	playdirGz := gziphandler.GzipHandler(http.HandlerFunc(playdir))
	readDirGz := gziphandler.GzipHandler(http.HandlerFunc(readDir))
	stopGz := gziphandler.GzipHandler(http.HandlerFunc(stop))

	// http.Handle("/", indexGz)
	http.Handle("/init", initialGz)
	http.Handle("/ls", readDirGz)
	http.Handle("/next", nextGz)
	http.Handle("/play", playGz)
	http.Handle("/playdir", playdirGz)
	http.Handle("/stop", stopGz)
	http.Handle("/up", cdUpGz)
}

func launchWindow() {
	cmd := exec.Command("electron", "/Users/emchang3/go/src/music_player/main.js")

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(stdout.String())
	fmt.Printf(stderr.String())
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routeHandler()
	launchWindow()

	port := getPort()
	fmt.Printf("\n--- Listening:%v\n\n", port)

	log.Fatal(http.ListenAndServe(port, nil))

	// cert, key := getCreds()

	// log.Fatal(http.ListenAndServeTLS(port, cert, key, nil))
}
