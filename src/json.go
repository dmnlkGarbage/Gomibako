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
	json, err := json.Marshal(message)

	if err != nil {
		fmt.Println(err)
		return
	}
	//標準出力に出してみる
	os.Stdout.Write(json)


	//デコードしてみる
	var decodeJson Message
	// json.Unmarshal undefined (type []byte has no field or method Unmarshal)
	json.Unmarshal(json, &decodeJson)

	if err != nil {
		fmt.Println(err)
		return
	}
	//print
	pp.Print(decodeJson)
}

