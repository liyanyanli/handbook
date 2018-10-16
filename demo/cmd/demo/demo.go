package main

import (
	"flag"
	"net/http"

	"fmt"
)

var hello string

func SayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Head:", req.Header)
	w.Write([]byte("Hello " + hello + "!\n"))
}

func main() {
	flag.StringVar(&hello, "city", "Hangzhou", "city name")

	flag.Parse()
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8000", nil)

}
