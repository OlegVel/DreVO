package main

import (
	"encoding/json"
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
		decoder := json.NewDecoder(request.Body)
		var jsonObject interface{}
		err := decoder.Decode(&jsonObject)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(jsonObject)
		writer.WriteHeader(http.StatusOK)
	}
}
