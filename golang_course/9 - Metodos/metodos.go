package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Save user data %s in database", u.name)
}

func main() {

	user1 := user{"Lucas", 26}
	fmt.Println(user1)
	user1.save()
}
