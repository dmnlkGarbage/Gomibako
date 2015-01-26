package main

import (
	"net/http"
	"time"
	"fmt"
)


func main() {
	http.HandleFunc("/", helloStreaming)
	http.ListenAndServe(":8000", nil)
}

func helloStreaming(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)

	for  {
		w.Write([]byte("hello\n"))
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(w, "hello end\n")
}
