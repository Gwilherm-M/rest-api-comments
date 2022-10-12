package controllers

import (
	"net/http"
	"rest-api-comments/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func HomePage(response http.ResponseWriter, request *http.Request) {
	message := bson.A{}
	message = append(
		message,
		"This REST server allows to manage comments linked to an AssetId.",
	)
	message = append(message, "Path: /comment/{id:[0-9a-z]+}")
	message = append(message, "METHODE: GET, POST, PUT, DELETE")
	message = append(message, "Path: /comments/{assetId}")
	message = append(message, "METHODE: GET")

	utils.OutResponse(
		response, http.StatusOK, "Welcome to the HomePage!", message,
	)
}
