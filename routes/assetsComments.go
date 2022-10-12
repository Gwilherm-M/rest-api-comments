package routes

import (
	"rest-api-comments/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CommentsRoute(router *mux.Router, clientDb *mongo.Client) {
	router.HandleFunc(
		"/assetsComments/{assetId}",
		controllers.HandlerComments(controllers.GetCommentsAssetId, clientDb),
	).Methods("GET")
	router.HandleFunc(
		"/assetsComments/",
		controllers.HandlerComments(controllers.GetAssetsId, clientDb),
	).Methods("GET")
}
