package utils

import (
	"HW1/handlers"

	"github.com/gorilla/mux"
)

func BuildProductResource(router *mux.Router, prefix string) {

	//GET /api/v1/item/{id} - возвращает информацию про товар с id и код 200 если такой товар в бд
	//сущестует. Если магазина с id в бд на данный момент нет - то сообщение "Error" : "Item with
	//that id not found" и код 404.
	router.HandleFunc(prefix+"/{id}", handlers.GetProductById).Methods("GET")

	//POST /api/v1/item/{id} - добавляет новый товар в бд (информацию из магазина считывать из
	//json). После добавления - возвращаем 201 , а также сообщение "Message" : "Item created". В
	//случае если товар с таким id уже существует - "Error" : "Ityem with that id already
	//exists", 400.
	router.HandleFunc(prefix, handlers.CreateProduct).Methods("POST")

	//PUT /api/v1/item/{id} - обновляет информацию про товар с id. Если такой товар имеется в бд -
	//выводим обновленный товар и код 202. В случае если товара с таким id нет в БД - "Error" :
	//"Item with that id not found" и код 404. (данный запрос реализовывать по желанию, т.к. на
	//лекции еще не успели полностью разобрать данный пункт)
	router.HandleFunc(prefix+"/{id}", handlers.UpdateProductById).Methods("PUT")

	//DELETE /api/v1/item/{id} - удаляет информацию про товар с id. Если такой товар имеется в БД -
	//выводим сообщение "Message": "Item deleted" и код 202. В противном случае - "Error" :
	//"Item with that id not found" и код 404. (данный запрос реализовывать по желанию, т.к. на
	//лекции еще не успели полностью разобрать данный пункт)
	router.HandleFunc(prefix+"/{id}", handlers.DeleteProductById).Methods("DELETE")
}

func BuildManyProductsResourcePrefix(router *mux.Router, prefix string) {
	//GET /api/v1/items - возвращает все имеющиеся на складе товары и код 200. В случае, если на
	//складе нет товаров в данный момент - вывести сообщение "Error" : "No one items found in
	//store back" и код 403
	router.HandleFunc(prefix, handlers.GetAllProducts).Methods("GET")
}
