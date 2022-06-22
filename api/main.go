package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	fmt.Println(config.Port)
	fmt.Println(config.ConnectionString)
	fmt.Println("rodando api")

	r := router.BuildRouter()

	log.Fatal(http.ListenAndServe(":5000", r))
}
