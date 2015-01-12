package main

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/k0kubun/pp"
)
type SearchResult struct {
	Events []struct {
	Event struct {
		Accepted       int    `json:"accepted"`
		Address        string `json:"address"`
		Catch          string `json:"catch"`
		Description    string `json:"description"`
		EndedAt        string `json:"ended_at"`
		EventID        int    `json:"event_id"`
		EventURL       string `json:"event_url"`
		Lat            string `json:"lat"`
		Limit          int    `json:"limit"`
		Lon            string `json:"lon"`
		OwnerID        int    `json:"owner_id"`
		OwnerNickname  string `json:"owner_nickname"`
		OwnerTwitterID string `json:"owner_twitter_id"`
		Place          string `json:"place"`
		StartedAt      string `json:"started_at"`
		Title          string `json:"title"`
		UpdatedAt      string `json:"updated_at"`
		URL            string `json:"url"`
		Waiting        int    `json:"waiting"`
	} `json:"event"`
} `json:"events"`
	ResultsReturned int `json:"results_returned"`
	ResultsStart    int `json:"results_start"`
}

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

