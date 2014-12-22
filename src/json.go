package main

import (
	"encoding/json"
	"fmt"

	"os"
)

type  Message struct {
	Id int64
	Data string
}
func main() {
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
}

