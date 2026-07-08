package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	op := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	numA, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	numB, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var result int

	switch op {
	case "add":
		result = numA + numB
	case "subtract":
		result = numA - numB
	case "multiply":
		result = numA * numB
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %d", result)
}

func main() {
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Server is running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}