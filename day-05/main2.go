package main

import "fmt"

type Notifier interface {
	Send()
}

type Email struct{}

func (e Email) Send() {
	fmt.Println("Sending email...")
}

type Sms struct{}

func (s Sms) Send() {
	fmt.Println("Sending SMS...")
}

func notify(n Notifier) {
	n.Send()
}

func main() {
	notify(Email{})
	notify(Sms{})
}
