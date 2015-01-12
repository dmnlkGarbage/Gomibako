package main

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/k0kubun/pp"
)


const (
	URL = "http://api.atnd.org/events/"
)
func main() {
	param := url.Values{}
	param.Add("keyword", "golang")
	param.Add("format", "json")

	resp, err := http.Get(URL + "?" + param.Encode())
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer resp.Body.Close()
	val, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return;
	}

	var result SearchResult
	if err = json.Unmarshal(val, &result);err != nil {
		fmt.Println(err)
		return
	}
	pp.Print(result)

}

