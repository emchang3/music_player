package main

import (
	"fmt"
	"log"
	"net/http"

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
	indexGz := gziphandler.GzipHandler(http.HandlerFunc(index))
	initialGz := gziphandler.GzipHandler(http.HandlerFunc(initial))
	nextGz := gziphandler.GzipHandler(http.HandlerFunc(next))
	playGz := gziphandler.GzipHandler(http.HandlerFunc(play))
	playdirGz := gziphandler.GzipHandler(http.HandlerFunc(playdir))
	readDirGz := gziphandler.GzipHandler(http.HandlerFunc(readDir))
	stopGz := gziphandler.GzipHandler(http.HandlerFunc(stop))

	http.Handle("/", indexGz)
	http.Handle("/init", initialGz)
	http.Handle("/ls", readDirGz)
	http.Handle("/next", nextGz)
	http.Handle("/play", playGz)
	http.Handle("/playdir", playdirGz)
	http.Handle("/stop", stopGz)
	http.Handle("/up", cdUpGz)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routeHandler()

	port := getPort()
	fmt.Printf("\n--- Listening:%v\n\n", port)

	log.Fatal(http.ListenAndServe(port, nil))

	// cert, key := getCreds()

	// log.Fatal(http.ListenAndServeTLS(port, cert, key, nil))
}
