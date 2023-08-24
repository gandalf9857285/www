package main

import (
	"fmt"
	"net/http"
)

func home_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Proverka")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ContaktЫ")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
