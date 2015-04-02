package main

import (
	"text/template"
	"os"
	"fmt"
)

type Member struct {
	id   int
	Name string
	Tech string
}

func main() {
	// template用
	const template_text = `template name is {{.Name}}`

	// Paraseする
	tpl := template.Must(template.New("mytemp").Parse(template_text))

	// 適当な構造体
	member := Member{1,"dmnlk", "Java"}

	// 埋め込んで標準出力
	if err := tpl.Execute(os.Stdout, member);err != nil {
		fmt.Println(err)
		return
	}
}
