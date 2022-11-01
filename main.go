package main

import (
	"fmt"
)

type userData struct {
	name  string
	email string
}

func main() {
	data := userData{name: "sajjad", email: "1234"}
	fmt.Println(data)
	fmt.Println(data.email)

	data.email = "email@mm.com"

	fmt.Println(data.email)
}
