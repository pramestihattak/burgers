package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":6969"
	fmt.Printf("Server started on port: %s\n", port)
	http.ListenAndServe(port, Router().Init())
}
