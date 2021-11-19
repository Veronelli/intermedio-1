package main

import "fmt"

//SMS email

type INotificacionFactory interface {
	SendNotificaction()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotificaction() {
	fmt.Println("Sending Notification via SMS")

}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}

}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"

}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"

}

//Email
type EmailNotificaction struct {
}

func (EmailNotificaction) SendNotificaction() {
	fmt.Println("Sending Notificacion via Email")

}

func (EmailNotificaction) GetSender() ISender {
	return EmailNotificactionSender{}

}

type EmailNotificactionSender struct {
}

func (EmailNotificactionSender) GetSenderMethod() string {
	return "Email"

}

func (EmailNotificactionSender) GetSenderChannel() string {
	return "Ses"

}

func getNotificationFactory(notificationType string) (INotificacionFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotificaction{}, nil
	}
	return nil, fmt.Errorf("No Notification tyoe")

}

func SendNotificaction(f INotificacionFactory) {
	f.SendNotificaction()
}

func getMethod(f INotificacionFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}
func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	SendNotificaction(smsFactory)
	SendNotificaction(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

}
