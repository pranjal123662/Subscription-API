package main

import (
	"fmt"
	"log"
	"net/http"
	"subscription/router"
)

func main() {
	r := router.Router()
	fmt.Println("server is running on port:7689")
	log.Fatal(http.ListenAndServe(":7689", r))
}
