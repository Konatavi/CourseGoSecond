package main

import (
	"log"
	"net/http"
	"os"

	"HW1/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                       string
	productResourcePrefix      string = apiPrefix + "/item"  //api/v1/item//
	manyProductsResourcePrefix string = apiPrefix + "/items" //api/v1/items/
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildProductResource(router, productResourcePrefix)
	utils.BuildManyProductsResourcePrefix(router, manyProductsResourcePrefix)

	log.Println("Router initalizing successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
