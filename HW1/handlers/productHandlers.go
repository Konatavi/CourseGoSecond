package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"HW1/models"

	"github.com/gorilla/mux"
)

//GET /api/v1/item/{id} - возвращает информацию про товар с id и код 200 если такой товар в бд
//сущестует. Если магазина с id в бд на данный момент нет - то сообщение "Error" : "Item with
//that id not found" и код 404.
func GetProductById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.MessageError{Error: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	product, ok := models.FindProductById(id)
	log.Println("Get product with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.MessageError{Error: "Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(product)
	}
}

//POST /api/v1/item/{id} - добавляет новый товар в бд (информацию из магазина считывать из
//json). После добавления - возвращаем 201 , а также сообщение "Message" : "Item created". В
//случае если товар с таким id уже существует - "Error" : "Ityem with that id already
//exists", 400.
func CreateProduct(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new product ....")
	var newProduct models.Product

	err := json.NewDecoder(request.Body).Decode(&newProduct)
	if err != nil {
		msg := models.MessageError{Error: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.MessageError{Error: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok := models.FindProductById(newProduct.ID)

	if ok {
		msg := models.MessageError{Error: "Item with that id already exists"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	models.DB = append(models.DB, newProduct)

	writer.WriteHeader(201)
	msg := models.Message{Message: "Item created"}
	json.NewEncoder(writer).Encode(msg)
}

//PUT /api/v1/item/{id} - обновляет информацию про товар с id. Если такой товар имеется в бд -
//выводим обновленный товар и код 202. В случае если товара с таким id нет в БД - "Error" :
//"Item with that id not found" и код 404. (данный запрос реализовывать по желанию, т.к. на
//лекции еще не успели полностью разобрать данный пункт)
func UpdateProductById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.MessageError{Error: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok := models.FindProductById(id)
	var newProduct models.Product
	if !ok {
		log.Println("item not found in data base")
		writer.WriteHeader(404)
		msg := models.MessageError{Error: "Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&newProduct)
	if err != nil {
		msg := models.MessageError{Error: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	for idSlice, val := range models.DB {
		if val.ID == id {
			models.DB[idSlice].Title = newProduct.Title
			models.DB[idSlice].Amount = newProduct.Amount
			models.DB[idSlice].Price = newProduct.Price
			newProduct = models.DB[idSlice]
			break
		}
	}

	writer.WriteHeader(202)
	json.NewEncoder(writer).Encode(newProduct)

}

//DELETE /api/v1/item/{id} - удаляет информацию про товар с id. Если такой товар имеется в БД -
//выводим сообщение "Message": "Item deleted" и код 202. В противном случае - "Error" :
//"Item with that id not found" и код 404. (данный запрос реализовывать по желанию, т.к. на
//лекции еще не успели полностью разобрать данный пункт)
func DeleteProductById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Deleting Product ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok := models.FindProductById(id)
	if !ok {
		log.Println("book not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.MessageError{Error: "Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//TODO: Нужно удалить book из DB - - DONE
	del_idSlice := -1
	for idSlice, val := range models.DB {
		if val.ID == id {
			del_idSlice = idSlice
			break
		}
	}

	models.DB = append(models.DB[0:del_idSlice], models.DB[del_idSlice+1:]...)
	writer.WriteHeader(202)
	msg := models.Message{Message: "Item deleted"}
	json.NewEncoder(writer).Encode(msg)
}
