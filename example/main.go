package main

import (
	"log"
	"net/http"
	"os"

	"github.com/srlightbody/gcsfileserver/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	s := gcsfileserver.server.Server{
		DirListPageSize: 100,
	}
	http.Handle("/", &s)
}
