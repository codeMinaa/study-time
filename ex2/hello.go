package main


import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w,"Hello ,%s!",name)
}
func main() {
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("server is running on :8080")

	http.ListenAndServe(":8080",nil)
}
