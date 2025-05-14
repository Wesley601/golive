package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

const DEFAULT_PORT = "5500"

func main() {
	p := flag.String("port", DEFAULT_PORT, "port to listen on")
	flag.Parse()

	if !isNumeric(*p) {
		fmt.Println("Invalid port number")
		return
	}
	port := ":" + *p

	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Printf("Starting to serve at http://localhost%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err.Error())
	}
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
