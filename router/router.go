package router

import (
	"github.com/ARMAAN199/practiceURL/controller"
	"github.com/gorilla/mux"
)

func UrlRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", controller.CreateShortUrl).Methods("POST")
	router.HandleFunc("/getOriginal/{short}", controller.GetActualUrl).Methods("GET")

	return router
}
