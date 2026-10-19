package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("PayFlow server 
	is starting...")

	http.HandleFunc("/payments", paymentsHandler)

	http.ListenAndServe(":8080", nil)
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to PayFlow!")
}

