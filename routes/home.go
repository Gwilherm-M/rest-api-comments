package routes

import (
	"rest-api-comments/controllers"

	"github.com/gorilla/mux"
)

func HomeRoute(router *mux.Router) {
	router.HandleFunc("/", controllers.HomePage)
}
