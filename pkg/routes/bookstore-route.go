package routes

import (
	"go-library-API-pr3/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisteredBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{ID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{ID}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{ID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
}
