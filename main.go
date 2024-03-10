package main

import (
	"fmt"
	"github.com/ARMAAN199/practiceURL/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go is on")
	err := http.ListenAndServe(":3008", router.UrlRouter())
	if err != nil {
		log.Fatal(err)
	}
}
