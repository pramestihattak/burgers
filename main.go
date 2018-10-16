package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":6969", ChiRouter().InitRouter())
}
