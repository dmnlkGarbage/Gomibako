package main

import "fmt"

// sturct
type UrlParam struct {
	Keyword string
	Format  string
}

func NewParam(keyword string) *UrlParam {
	u := new(UrlParam)
	u.Format = "json"
	u.Keyword = keyword
	return u
}

func NewParam2(keyword string) *UrlParam {
	return &UrlParam{keyword, "json"}
}

func makeParam(keyword string) UrlParam {
	return UrlParam{keyword, "json"}
}

func main() {
	param := UrlParam{}
	param.Keyword = "golang"
	param.Format = "json"
	fmt.Println(param)

	p1 := NewParam("golang")
	fmt.Println(p1)

	p2 := NewParam2("golang")
	fmt.Println(p2)

	p3 := makeParam("golang")
	fmt.Println(p3)

}
