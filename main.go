package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/postgres-go/router"
)

func main(){
	r := router.Router()
	fmt.Println("Starting Server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}