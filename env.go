package main

import (
	"fmt"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8080"
	}

	return fmt.Sprintf(":%v", port)
}

func getToken() string {
	token := os.Getenv("TOKEN")

	if token == "" {
		return "nah, bruh."
	}

	return token
}

func getCreds() (string, string) {
	cert := os.Getenv("CERT")
	key := os.Getenv("KEY")

	return cert, key
}
