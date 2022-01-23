package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	http.ServeFile(w, r, "./static/index.html")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/minha-api", handlerOla)
	fmt.Println("Server is up and listening on port 8081.")
	http.ListenAndServe(":8081", nil)
}

func handlerOla(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	url := r.Host

	w.WriteHeader(http.StatusOK)

	resp := make(map[string]string)
	resp["message"] = "Status Created"
	resp["api-url"] = string(url)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	//http.ServeFile(w, r, "./static/index.html")
}
