package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pavanilla/middleware"
	"github.com/pavanilla/router"
)

func main() {
	r := router.Router()
	middleware.Init()
	fmt.Println("server is started on the port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
