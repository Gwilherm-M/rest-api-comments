package main

import (
	"log"
	"net/http"
	"rest-api-comments/routes"
	"rest-api-comments/utils"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("[INFO]\tLaunch serveur ...")

	router := mux.NewRouter()
	router.Use(utils.JsonContentType)
	dataBase, cancel := utils.ConnectDb()
	defer cancel()

	routes.HomeRoute(router)
	routes.CommentRoute(router, dataBase)
	routes.CommentsRoute(router, dataBase)

	log.Print("[INFO]\tServer is ready")
	log.Fatalln("[ERROR]\t", http.ListenAndServe("localhost:8000", router))
}
