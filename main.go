package main

import (
	"fmt"
	"net/http"
)

func main() {
	msg := hello()
	adios := adios()
	fmt.Println(msg)
	fmt.Println("*******")
	fmt.Println(adios)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Captain Canary's Solution Architecture Demo Go App!")
	})

	http.ListenAndServe(":8090", nil)
}
