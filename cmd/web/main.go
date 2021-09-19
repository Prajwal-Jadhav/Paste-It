package main

import (
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte("This is the home page"))
	})

	err := http.ListenAndServe(":4000", serveMux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
