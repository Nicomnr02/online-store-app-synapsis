package main

import (
	"log"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	}))

	if err != nil {
		log.Fatal(err)
	}
}
