package main

import (
	"fmt"
	"os"

	"github.com/rem7/goprowl"
	"strconv"
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
		Application: "Foo",
		Description: "Foobar!",
		Event:       "Bar",
		Priority:    "1",
		Providerkey: "",
		Url:         "www.foobar.com",
	}

	c := make(chan int)
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		n.Description = "test" + strconv.Itoa(i)
		go func() {
			err = p.Push(n)
			c <- 1
		}()
	}

	fmt.Println(<-c)
	if err != nil {
		fmt.Println("Unable to send Prowl notification! - " + err.Error())
	}
}
