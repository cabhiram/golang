package main

import "fmt"

type User struct {
	name        string
	phoneNumber string
}

func (u *User) Modify(name string) {
	(*u).name = name
}

func (u User) Print() {

	fmt.Println(u)
}
