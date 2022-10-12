package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"rest-api-comments/models"
	"rest-api-comments/utils"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetComment(
	response http.ResponseWriter,
	request *http.Request,
	db *mongo.Collection,
	ctx context.Context,
) {
	params := mux.Vars(request)
	commentId := params["id"]
	var comment models.Comment

	objId, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		utils.OutResponse(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = db.FindOne(ctx, models.Comment{Id: objId}).
		Decode(&comment)
	if err != nil {
		utils.OutResponse(response, http.StatusNotFound, err.Error(), nil)
		return
	}

	utils.OutResponse(response, http.StatusOK, "", comment)
}

func CreateComment(
	response http.ResponseWriter,
	request *http.Request,
	db *mongo.Collection,
	ctx context.Context,
) {
	var comment models.Comment

	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		utils.OutResponse(response, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if comment.AssetId == "" {
		utils.OutResponse(
			response, http.StatusBadRequest, "Please set assetId.", nil,
		)
		return
	}
	if comment.Text == "" {
		utils.OutResponse(
			response, http.StatusBadRequest, "Please set text.", nil,
		)
		return
	}

	comment.CreateAt = primitive.NewDateTimeFromTime(time.Now())
	comment.LastUpdateAt = primitive.NewDateTimeFromTime(time.Now())

	result, err := db.InsertOne(ctx, comment)
	if err != nil {
		utils.OutResponse(
			response, http.StatusInternalServerError, err.Error(), nil,
		)
		return
	}
	utils.OutResponse(response, http.StatusOK, "", result)
}

func UpdateComment(
	response http.ResponseWriter,
	request *http.Request,
	db *mongo.Collection,
	ctx context.Context,
) {
	params := mux.Vars(request)
	commentId := params["id"]
	var comment models.Comment

	objId, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		utils.OutResponse(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		utils.OutResponse(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	comment.Id = objId
	filter := bson.M{"_id": bson.M{"$eq": objId}}
	update := bson.M{"$set": bson.D{
		{Key: "text", Value: comment.Text},
		{Key: "lastUpdate", Value: time.Now()},
	}}

	result, err := db.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.OutResponse(
			response, http.StatusInternalServerError, err.Error(), nil,
		)
		return
	}

	if result.MatchedCount == 1 {
		utils.OutResponse(
			response, http.StatusAccepted, "Comment is updated", comment,
		)
	} else {
		utils.OutResponse(
			response, http.StatusNotFound, "Comment not found", nil,
		)
	}
}

func DeleteComment(
	response http.ResponseWriter,
	request *http.Request,
	db *mongo.Collection,
	ctx context.Context,
) {
	params := mux.Vars(request)
	commentId := params["id"]
	var comment models.Comment

	objId, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		utils.OutResponse(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	comment.Id = objId
	result, err := db.DeleteOne(ctx, comment)
	if err != nil {
		utils.OutResponse(
			response, http.StatusInternalServerError, err.Error(), nil,
		)
		return
	}

	if result.DeletedCount == 1 {
		utils.OutResponse(
			response, http.StatusAccepted, "Comment is removed.", comment,
		)
	} else {
		utils.OutResponse(
			response, http.StatusNotFound, "Comment not found.", nil,
		)
	}
}

func HandlerComment(
	objFunc models.RequestCollectType,
	ClientDb *mongo.Client,
) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		db := utils.GetCommentCollection(ClientDb)
		objFunc(response, request, db, ctx)
	}
}
