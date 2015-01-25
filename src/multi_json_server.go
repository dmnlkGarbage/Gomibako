package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", streaming)
	http.ListenAndServe(":8000", nil)
}

func streaming(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	for i := 0; i < 5; i++ {
		w.Write([]byte("hello¥n"))
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}
}
