package main


import "fmt"

func CreateNotification(notificationType string) (Notification, error) {
	if notificationType == "email" {
		return EmailNotification{}, nil
	} else if notificationType == "sms" {
		return SMSNotification{}, nil
	}
	return nil, fmt.Errorf("invalid notification type")
}