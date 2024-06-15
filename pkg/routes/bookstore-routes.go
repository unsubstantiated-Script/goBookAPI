package routes

import (
	"github.com/gorilla/mux"
	"goBookAPI/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
}
