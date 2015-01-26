package main

import (
	"flag"
	"fmt"
)

// flagパッケージ試す
func main() {
	fmt.Println("flag test")

	//コマンドライン引数からhogeというのをもってくる。第2引数はデフォルト値、第3は使い方説明的な
	//戻り値はstringのポインタ
	var a *string = flag.String("hoge", "c", "usage")
	//パースしないと使えない
	flag.Parse()
	fmt.Println(*a)

	//$ go run flag_sample.go                                                                                                                    [18:48:38]
	//flag test
	//c

	//$ go run flag_sample.go --hoge=a                                                                                                           [18:48:42]
	//flag test
	//a

	//$ go run flag_sample.go -hoge a                                                                                                            [18:49:34]
	//flag test
	//a
}
