package routes

import (
	"rest-api-comments/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CommentRoute(router *mux.Router, clientDb *mongo.Client) {
	router.HandleFunc(
		"/comment/{id:[0-9a-z]+}",
		controllers.HandlerComment(controllers.GetComment, clientDb),
	).Methods("GET")
	router.HandleFunc(
		"/comment/{id:[0-9a-z]+}",
		controllers.HandlerComment(controllers.DeleteComment, clientDb),
	).Methods("DELETE")
	router.HandleFunc(
		"/comment/{id:[0-9a-z]+}",
		controllers.HandlerComment(controllers.UpdateComment, clientDb),
	).Methods("PUT")
	router.HandleFunc(
		"/comment",
		controllers.HandlerComment(controllers.CreateComment, clientDb),
	).Methods("POST")
}
