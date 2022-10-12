package utils

import (
	"encoding/json"
	"net/http"
	"rest-api-comments/models"
)

func OutResponse(
	response http.ResponseWriter, status int, text string, data interface{},
) error {
	response.WriteHeader(status)
	result := models.NewResponse(status, text, data)
	err := json.NewEncoder(response).Encode(result)

	return err
}
