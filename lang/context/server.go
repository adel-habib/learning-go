package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle started")
	defer log.Printf("Handler ended")
	ctx := r.Context()
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "HELLO FROM SERVER.")
	case <-ctx.Done():
		log.Printf(ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)

	}
}
