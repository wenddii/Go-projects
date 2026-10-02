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
	fmt.Println("Sending sms...")
}
func main() {
	var n Notifier = Email{}
	n.Send()

	n = Sms{}
	n.Send()

}
