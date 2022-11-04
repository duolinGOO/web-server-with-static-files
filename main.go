package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellofunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprint(w, "the method is not GET")
	}
	if r.URL.Path != "/hello" {
		return
	}
	fmt.Fprintf(w, "As salam aleykum va rahmatullah")
}
func formhandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "mistake")
		return
	}
	fmt.Fprintf(w, "success\n")
	name := r.FormValue("name1")
	address := r.FormValue("address1")
	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v", address)
}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellofunc)
	http.HandleFunc("/form", formhandle)
	if err := http.ListenAndServe(":4545", nil); err != nil {
		log.Fatal("error")
	}

}
