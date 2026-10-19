package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /members", membersHandler)
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("POST /payments", createPaymentHandler)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func membersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Here is the list of members")
}

func createPaymentHandler(w http.ResponseWriter, r *http.Request) {
	amount := r.URL.Query().Get("amount")

	if amount == "" {
		http.Error(
			w,
			"Payment amount is required",
			http.StatusBadRequest,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Payment created successfully")
}

