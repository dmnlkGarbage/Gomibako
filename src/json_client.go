package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/k0kubun/pp"
)

func main() {

	res, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(res.Body)
	for {
		if ok := scanner.Scan(); !ok {
			break
		}

		//まずinterfaceでうける
		//json化
		var result interface{}

		if err = json.Unmarshal(scanner.Bytes(), &result); err != nil {
			fmt.Println(err)
			continue
		}
		msg := result.(map[string]interface{})
		if _, ok := msg["Title"]; ok {
			u := new(User)
			if err = json.Unmarshal(scanner.Bytes(), &u); err != nil {
				break
			}
			pp.Print(u)
		}
		if _, ok := msg["Name"]; ok  {
			e := new(Event)
			if err = json.Unmarshal(scanner.Bytes(), &e); err != nil {
				break
			}
			pp.Print(e)
		}

	}
}

type User struct {
	Name string
	Age  int64
}

type Event struct {
	Title string
}
