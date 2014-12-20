package main

import (
	"fmt"
)

type App struct {
	appId int
	appName string
}

// override的なことは出来ない？
func (a App) changeAppId(id int) {
	a.appId = id
	fmt.Println(a.appId)
}
