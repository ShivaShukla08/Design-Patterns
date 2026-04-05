package main
import "fmt"

type Notification interface{
	SendNotification() 
}


type EmailNotification struct{}

func (e EmailNotification) SendNotification() {
	fmt.Println("Sending Email Notification")
}

type SMSNotification struct{}

func (s SMSNotification) SendNotification() {
	fmt.Println("Sending SMS Notification")
}