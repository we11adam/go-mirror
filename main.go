package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9080"
	}

	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	details := make(map[string]interface{})
	details["method"] = r.Method
	details["path"] = r.URL.Path
	details["query"] = r.URL.Query()
	details["headers"] = r.Header

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		details["body"] = string(body)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(details)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
