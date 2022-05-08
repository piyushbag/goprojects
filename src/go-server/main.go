package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", fileServer)
	http.HandleFunc("/form", form)
	http.HandleFunc("/hello", hello)

	fmt.Printf("Starting server ar port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
