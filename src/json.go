package main

import (
	"encoding/json"
	"fmt"

	"os"
	"github.com/k0kubun/pp"
)

type  Message struct {
	Id int64
	Data string
}
func main() {

	//エンコーディング試す
	//構造体つくる
	message := Message{2, "data"}
	// エンコーディング。バイト列
	enc, err := json.Marshal(message)

	if err != nil {
		fmt.Println(err)
		return
	}
	//標準出力に出してみる
	os.Stdout.Write(enc)


	//デコードしてみる
	var decodeJson Message
	// jsonとかいう変数が悪い
	json.Unmarshal(enc, &decodeJson)

	if err != nil {
		fmt.Println(err)
		return
	}
	//print
	pp.Print(decodeJson)
}

