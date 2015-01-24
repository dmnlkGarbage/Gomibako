package main

import "fmt"

// sturct
type User struct {
	Age int64
	Name string
}

func NewUser (name string) * User {
	u := new(User)
	u.Name = name
	u.Age = 1000
	return u
}

func main() {
	// コンストラクタ
	user := User{0, "a"}
	fmt.Println(user)
	// 代入
	user.Age = 10
	fmt.Println(user)

	u := NewUser("dmnlk")
	fmt.Println(u)
}


