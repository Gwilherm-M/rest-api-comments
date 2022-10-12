package models

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AssetId      string             `json:"assetId,omitempty" bson:"assetId,omitempty"`
	Text         string             `json:"text,omitempty" bson:"text,omitempty"`
	CreateAt     primitive.DateTime `json:"createAt,omitempty" bson:"createAt,omitempty"`
	LastUpdateAt primitive.DateTime `json:"lastUpdate,omitempty" bson:"lastUpdate,omitempty"`
}

type RequestCollectType func(
	w http.ResponseWriter,
	r *http.Request,
	db *mongo.Collection,
	ctx context.Context,
)

type Response struct {
	Status int         `json:"status"`
	Text   string      `json:"text"`
	Data   interface{} `json:"data"`
}

func NewResponse(status int, text string, data interface{}) *Response {
	return &Response{
		Status: status,
		Text:   text,
		Data:   data,
	}
}
