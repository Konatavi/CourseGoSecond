package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"HW1/models"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

//GET /api/v1/items - возвращает все имеющиеся на складе товары и код 200. В случае, если на
//складе нет товаров в данный момент - вывести сообщение "Error" : "No one items found in
//store back" и код 403
func GetAllProducts(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	if len(models.DB) == 0 {
		writer.WriteHeader(403)
		msg := models.MessageError{Error: "No one items found in store back"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	log.Println("Get infos about all items in database")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
