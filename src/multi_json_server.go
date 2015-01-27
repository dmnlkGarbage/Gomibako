package main

import (
	"net/http"
	"time"
	"fmt"
	"math/rand"
	"encoding/json"
)


func main() {
	http.HandleFunc("/", helloStreaming)
	http.ListenAndServe(":8000", nil)
}

func helloStreaming(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)

	for  {
		rand.Seed(time.Now().Unix())
		rand := rand.Intn(2)
		if rand == 0 {
			u := new(User)
			u.Name = "a"
			u.Age = 22
			enc, _ := json.Marshal(u)
			w.Write(enc)
		} else {
			e := new(Event)
			e.Title = "title"
			enc, _ := json.Marshal(e)
			w.Write(enc)
		}

		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(w, "hello end\n")
}

type User struct {
	Name string
	Age int64
}

type Event struct {
	Title string
}
