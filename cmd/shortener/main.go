package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	full_url := r.URL
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Location", full_url.String())
		w.WriteHeader(http.StatusOK)
	case "POST":
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		data := map[string]string{"Short-url": full_url.String()}
		json.NewEncoder(w).Encode(data)
	}
}

func main() {

	server := &http.Server{
		Addr: "localhost:8080"}

	http.HandleFunc("/", GetHandler)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server Shutdown")
	}
}
