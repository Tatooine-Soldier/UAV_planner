package main

import (
	"fmt"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the homepage")
	checkErr(err)
}
func getHello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	// http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	fs := http.FileServer(http.Dir("../../dist"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":3333", nil)
	fmt.Printf("%v", err)
}
