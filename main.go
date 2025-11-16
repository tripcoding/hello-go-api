package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var version = "development"

type Greeting struct {
	Message  string `json:"message"`
	Language string `json:"language"`
}

var greetingsMap = map[string]string{
	"en": "Hello, World!",
	"pt": "Olá, Mundo!",
	"es": "¡Hola, Mundo!",
	"ru": "Привет, Мир!",
	"ja": "こんにちは、世界！",
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	lang := r.URL.Query().Get("lang")

	message, ok := greetingsMap[lang]

	if !ok {
		lang = "en"
		message = greetingsMap[lang]
	}

	resp := Greeting{
		Message:  message,
		Language: lang,
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		log.Printf("Error encoding json: %v", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Printf("Starting server on :8080 (version: %s)", version)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
