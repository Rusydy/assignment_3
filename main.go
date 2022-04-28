package main

import (
	"assignment_3/functions"
	"fmt"
	"net/http"
)

func main() {
	go functions.CreateJson()
	http.HandleFunc("/", functions.ReloadWeb)
	fmt.Println("Running in port 8080")
	http.ListenAndServe(":8080", nil)
}
