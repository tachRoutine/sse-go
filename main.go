package main

import (
	"fmt"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./"))
	mux.Handle("/", fs)
	mux.HandleFunc("/events", sseHandler)

	fmt.Println("Listening on http://localhost:8080")
	if err:= http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}