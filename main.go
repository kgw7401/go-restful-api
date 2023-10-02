package main

import (
	"net/http"

	"github.com/go-web/go-restful-api/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}