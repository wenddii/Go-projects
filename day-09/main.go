package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/members", membersHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func membersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Here is the list of members")
}
