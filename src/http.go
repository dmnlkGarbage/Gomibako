package main

import (
	"net/url"
	"net/http"
	"fmt"
	"github.com/k0kubun/pp"
)

var URL = "http://api.atnd.org/events/"


func main() {
	values := url.Values{}
	values.Add("keyword", "golang")
	values.Add("format", "json")
	fmt.Println(URL + "?" + values.Encode())
	resp, err := http.Get(URL + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	pp.Print(resp)
	fmt.Printf("%+v", resp.Body)

}
