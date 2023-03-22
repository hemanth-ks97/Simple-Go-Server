package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST successful\n")
	name := r.FormValue("name")
	phone := r.FormValue("phone")

	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Phone: %v\n", phone)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	
	if r.Method != "GET"{
		http.Error(w, "Method not supported", http.StatusNotFound)
		return 
	}

	fmt.Fprintf(w, "HELLO!!")
}


func main(){
	fs := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fs)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}
}