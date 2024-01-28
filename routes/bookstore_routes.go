package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/gowdaganesh005/MySQLserver/controllers"
)

var RegisterBookStoreRoutes = func(router *chi.Mux) {
	router.Get("/Book/", controllers.GetBook)
	router.Get("/Book/{id}", controllers.GetBookbyId)
	router.Post("/Book/", controllers.CreateBook)
	router.Post("/Book/{id}", controllers.UpdateBook)
	router.Delete("/Book/{id}", controllers.DeleteBook)

}
