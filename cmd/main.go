package main

import (
	"Morse-Code-Converter/internal/server"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "MORSE: ", log.LstdFlags)
	srv := server.NewServer(l)

	l.Println("Starting server on :8080")

	if err := srv.Server.ListenAndServe(); err != nil {
		l.Fatal(err)
	}
}
