package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseJson(r *http.Request, payload interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("could not read the request:", err)
	}
	err = json.Unmarshal(body, payload)
	if err != nil {
		log.Println("Could not un marshal the request:", err)
	}
}
