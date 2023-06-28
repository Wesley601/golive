package main

import (
	"fmt"
	"net/http"
)

const PORT = ":5500"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Printf("Start to serve at http://localhost%s\n", PORT)
	http.ListenAndServe(PORT, nil)
}
