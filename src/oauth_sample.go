package main

import (
	"github.com/mrjones/oauth"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/k0kubun/pp"
	"flag"
)

//OauthClient
type Client struct {
	consumer *oauth.Consumer
	accessToken *oauth.AccessToken
}

func main() {
	client := initialize()
	result, err := client.Get("https://api.twitter.com/1.1/statuses/home_timeline.json", nil)
	if err != nil {
		fmt.Println(err)
	}
	pp.Print(result)
}

// client初期化
func initialize() *Client {

	//それぞれのkeyを実行時引数から取得
	var consumerKey *string = flag.String(
		"consumerkey",
		"",
		"Consumer Key from Twitter. See: https://dev.twitter.com/apps/new")

	var consumerSecret *string = flag.String(
		"consumersecret",
		"",
		"Consumer Secret from Twitter. See: https://dev.twitter.com/apps/new")

	var accessToken *string = flag.String(
		"accessToken",
		"",
		"Consumer Key from Twitter. See: https://dev.twitter.com/apps/new")

	var accessTokenSecret *string = flag.String(
		"accesssecret",
		"",
		"Consumer Secret from Twitter. See: https://dev.twitter.com/apps/new")

	//parseしないと取れない
	flag.Parse()
	fmt.Println(*consumerKey)
	fmt.Println(*consumerSecret)
	fmt.Println(*accessToken)
	fmt.Println(*accessTokenSecret)


	//初期化
	client := new(Client)
	//oauthパッケージでtwitterにconsumerkeyを元に取得
	client.consumer = oauth.NewConsumer(*consumerKey, *consumerSecret, oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
	//accessTokenを渡す。
	//ユーザーに認証させることもできるが面倒なので
	client.accessToken = &oauth.AccessToken{*accessToken, *accessTokenSecret, nil}
	return client
}

// clientを使ってTwitterからデータ取得する
func (client *Client) Get(url string, params map[string]string) (interface {}, error) {
	//Twitterにget
	response, err := client.consumer.Get(url, params, client.accessToken)
	if err != nil {
		fmt.Errorf("error")
		return nil, err
	}
	//close大事
	defer response.Body.Close()

	//jsonっぽいものにする
	bin, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("error")
		return nil, err
	}
	//unMarshalして返す。interfaceで受けるとstruct定義しなくていいのでちょっと楽
	var result interface{}
	err = json.Unmarshal(bin, &result)
	return result, err
}

