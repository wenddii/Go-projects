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

today i learned interfaces they are like the implements keyword in java and it

there is an interface called Notifier and everything that holds a Send() method is notifier 

then we created a struct called email and a method/function that recieves a variable called e and which is a type of email and that prints Sending email... and since it has Send() method in it. is a notifier same thing for SMS 

then we went over to the main function where everything happens 

and we instantiated a variable called n and it is a notifier and it type is Email 

then when we do n.Send() it prints the sending email 

and we go over and changed n to the Sms type and we called n.Send() and it prints Sending sms 



























