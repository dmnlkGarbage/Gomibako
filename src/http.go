package main

import (
	"net/url"
	"net/http"
	"fmt"
//	"github.com/k0kubun/pp"
	"io/ioutil"
	"github.com/k0kubun/pp"
)

var URL = "http://api.atnd.org/events/"


func main() {

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

}
