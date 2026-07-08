package main

import (
	"fmt"
	"net/http"
)

func pinghandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"pong")
}
func main() {
	http.HandleFunc("/ping", pinghandler)

	fmt.Println("server is running on :8080")
	
	http.ListenAndServe(":8080",nil)
}