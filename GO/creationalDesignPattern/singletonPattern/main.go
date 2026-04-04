package main

import "fmt"

func main() {
	fmt.Println("Start")

	service := NewUserService()
	service.GetUsers()

	db1 := GetInstance()
	db2 := GetInstance()

	fmt.Println(db1 == db2) 
}