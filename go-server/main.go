package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(res, "Hello!")
}

func fromhandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm err: %v", err)
	}
	fmt.Fprint(res, "Post request successfull")
	name := req.FormValue("name")
	email := req.FormValue("email")
	password := req.FormValue("password")
	fmt.Fprintf(res, "Name = %s\nEmail = %s\nPassword = %s\n", name, email, password)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", fromhandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("server starting at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
