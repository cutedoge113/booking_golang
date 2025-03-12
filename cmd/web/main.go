package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}
