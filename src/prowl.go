package main

import (
	"github.com/rem7/goprowl"
	"fmt"
	"os"
)


func main() {
	key := os.Getenv("prowl_key")
	var p goprowl.Goprowl
	err := p.RegisterKey(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	n := &goprowl.Notification{
		Application : "Foo",
		Description: "Foobar!",
		Event : "Bar",
		Priority : "1",
		Providerkey : "",
		Url: "www.foobar.com",
	}

	go func () {
		err = p.Push(n)
	}()

	if err != nil {
		fmt.Println("Unable to send Prowl notification! - " + err.Error())
	}
}

