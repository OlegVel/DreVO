package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", SiteHandler())
	if err := http.ListenAndServe(":0050", nil); err != nil {
		log.Fatal(err)
	}
}

func SiteHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}
