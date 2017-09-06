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

	indexGz := gziphandler.GzipHandler(http.HandlerFunc(index))
	initialGz := gziphandler.GzipHandler(http.HandlerFunc(initial))
	readDirGz := gziphandler.GzipHandler(http.HandlerFunc(readDir))
	cdUpGz := gziphandler.GzipHandler(http.HandlerFunc(cdUp))
	playGz := gziphandler.GzipHandler(http.HandlerFunc(play))
	pauseGz := gziphandler.GzipHandler(http.HandlerFunc(pause))
	contGz := gziphandler.GzipHandler(http.HandlerFunc(cont))
	nextGz := gziphandler.GzipHandler(http.HandlerFunc(next))

	http.Handle("/", indexGz)
	http.Handle("/init", initialGz)
	http.Handle("/ls", readDirGz)
	http.Handle("/up", cdUpGz)
	http.Handle("/play", playGz)
	http.Handle("/pause", pauseGz)
	http.Handle("/cont", contGz)
	http.Handle("/next", nextGz)
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
