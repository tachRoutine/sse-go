package main

import (
	"fmt"
	"net/http"
	"time"
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
	
	for i := range 10 {
		fmt.Fprintf(w, "data: Message %d\n\n", i)
		fmt.Println("Sent message", i)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		time.Sleep(1 * time.Second)
	}
}