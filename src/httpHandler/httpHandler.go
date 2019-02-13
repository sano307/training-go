package main

import (
	"fmt"
	"log"
	"net/http"
)

func name(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Inseo Kim")
}

func company(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Cyberagent")
}

func main() {
	http.HandleFunc("/name", name)
	http.HandleFunc("/company", company)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
