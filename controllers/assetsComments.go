package controllers

import (
	"context"
	"fmt"
	"net/http"
	"rest-api-comments/models"
	"rest-api-comments/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCommentsAssetId(
	response http.ResponseWriter,
	request *http.Request,
	db *mongo.Collection,
	ctx context.Context,
) {
	params := mux.Vars(request)
	query := request.URL.Query()
	qLimit := query["limit"]
	qSkip := query["skip"]
	qCreateAt := query["createAt"]
	qUpdateAt := query["updateAt"]

	var bFilter primitive.D
	if len(qCreateAt) > 0 {
		switch qCreateAt[0] {
		case "asc":
			bFilter = append(bFilter, bson.E{Key: "createAt", Value: 1})
		case "des":
			bFilter = append(bFilter, bson.E{Key: "createAt", Value: -1})
		}
	}

	if len(qUpdateAt) > 0 {
		switch qUpdateAt[0] {
		case "asc":
			bFilter = append(bFilter, bson.E{Key: "lastUpdate", Value: 1})
		case "des":
			bFilter = append(bFilter, bson.E{Key: "lastUpdate", Value: -1})
		}
	}

	if len(bFilter) == 0 {
		bFilter = append(bFilter, bson.E{Key: "createAt", Value: -1})
	}

	opts := options.Find().SetSort(bFilter)
	if len(qLimit) > 0 {
		limit, err := strconv.ParseInt(qLimit[0], 10, 64)
		if err == nil {
			opts.SetLimit(limit)
		}
	}
	if len(qSkip) > 0 {
		skip, err := strconv.ParseInt(qSkip[0], 10, 64)
		if err == nil {
			opts.SetSkip(skip)
		}
	}

	assetId := params["assetId"]
	cursor, err := db.Find(ctx, bson.M{"assetId": assetId}, opts)
	if err != nil {
		defer cursor.Close(ctx)
		utils.OutResponse(
			response, http.StatusInternalServerError, err.Error(), nil,
		)
		return
	}

	var assetComments []models.Comment
	if err = cursor.All(ctx, &assetComments); err != nil {
		utils.OutResponse(
			response, http.StatusInternalServerError, err.Error(), nil,
		)
		return
	}

	if len(assetComments) >= 1 {
		utils.OutResponse(
			response, http.StatusAccepted, "", assetComments,
		)
	} else {
		message := fmt.Sprintf("No comments found for AssetId %s", assetId)
		utils.OutResponse(
			response, http.StatusNotFound, message, nil,
		)
	}
}

func GetAssetsId(
	response http.ResponseWriter,
	request *http.Request,
	db *mongo.Collection,
	ctx context.Context,
) {
	assetsId, err := db.Distinct(ctx, "assetId", bson.M{})
	if err != nil {
		utils.OutResponse(response, http.StatusNotFound, err.Error(), nil)
		return
	}

	if len(assetsId) == 0 {
		utils.OutResponse(
			response, http.StatusNotFound, "Not assetsId.", nil,
		)
		return
	}
	utils.OutResponse(response, http.StatusOK, "", assetsId)
}

func HandlerComments(
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
