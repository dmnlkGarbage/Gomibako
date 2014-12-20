package main

import (
	"fmt"
)

func main() {
	var myApp = App{2, "aaaa"}
	fmt.Println(myApp.appId)

	//このスコープのmyAppの中身は変えない
	myApp.changeAppId(13)
	fmt.Println(myApp.appId)


	//中身を変更できる
	myApp.changeAppIdPointer(15)
	fmt.Println(myApp.appId)

}


// scan.goに定義されてるAppっていう構造体にメソッドを生やす。


// appのidを値渡しで変える。インスタンス自体の中身は変えない
func (a App) changeAppId(id int) {
	a.appId = id
	fmt.Println(a.appId)
}


// appのidを参照渡しで変える。インスタンス自体の中身を変更できる。ポインタ使ってる
func (a *App) changeAppIdPointer(id int) {
	a.appId = id
	fmt.Println(a.appId)
}
