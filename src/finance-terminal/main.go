package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"finance-terminal/backend"
)

func main() {
	httpPort := 8081
	log.Printf("Starting Server on port %v\n", httpPort)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		assets := r.Form["assets"]
		debts := r.Form["debts"]
		http.Redirect(w, r, "/", http.StatusSeeOther)
		backend.Terminal(assets[0], debts[0])

	}
}
