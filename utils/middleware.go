package utils

import "net/http"

func JsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Add(
				"content-type", "application/json; charset=UTF-8",
			)
			next.ServeHTTP(response, request)
		})
}
