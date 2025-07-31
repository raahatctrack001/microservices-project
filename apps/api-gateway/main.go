package main

import (
	"api-gateway/routers"
	"api-gateway/utils"
	"log"
	"net/http"
)

func main() {
	utils.LoadEnv()

	handler := router.NewRouter()

	log.Println("API Gateway running at :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
