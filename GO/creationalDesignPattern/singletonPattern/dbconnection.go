package main

import (
	"fmt"
	"sync"
)

type DatabaseConnection struct {
	connection string
}

var instance *DatabaseConnection
var lock sync.Mutex

func initialize() {
	fmt.Println("Creating DB Connection")
	instance = &DatabaseConnection{
		connection: "Dummy DB",
	}
}

func GetInstance() *DatabaseConnection {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil{
		initialize()
	}
	return  instance
}

func (db *DatabaseConnection) ExecuteQuery(query string) {
	fmt.Println("Executing", query)
}