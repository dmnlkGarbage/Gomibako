package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/k0kubun/pp"
)

var URL = "http://api.atnd.org/events/"

func main() {

	//普通にhttpパッケージ使ってみる

	//パラメータの組み立て
	values := url.Values{}
	values.Add("keyword", "golang")
	values.Add("format", "json")

	// Encodeすると?hoge=fuga&piyo=poyo みたいにしてくれる
	// getRequest送る
	resp, err := http.Get(URL + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}
	// ちゃんとclose
	defer resp.Body.Close()
	// レスポンスヘッダしか読めない？
	fmt.Println(resp)

	// ioutil.readallでバイト配列にする
	val, err := ioutil.ReadAll(resp.Body)

	// バイト配列を文字列にして表示する
	pp.Print(string(val))

	// httpClient使ってみる
	req, nil := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := new(http.Client)
	resp1, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp1.Body.Close()
	fmt.Println(resp1)
	val1, err := ioutil.ReadAll(resp1.Body)
	pp.Print(string(val1))

}
