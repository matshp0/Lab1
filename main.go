package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
	})

	port := ":8795"
	fmt.Printf("Server running on http://localhost%s/time\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
