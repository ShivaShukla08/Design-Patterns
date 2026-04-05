package main

import "fmt"

func main() {
	notificationType := "sms"

	notification, err := CreateNotification(notificationType)
	if err != nil {
		fmt.Println(err)
		return
	}

	notification.SendNotification()
}