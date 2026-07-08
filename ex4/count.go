package main


import (
	"fmt"
	"io"
	"net/http"
	
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Send a POST request with text to count words")
	case http.MethodPost:
		body , err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		
		
		count := len(body)
	    fmt.Fprintf(w,"%d",count)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	
	
}
func main() {
	http.HandleFunc("/count",countHandler)

	fmt.Println("server running on :8080")

	if err := http.ListenAndServe(":8080",nil);err != nil {
		fmt.Println(err)

	}
	

}