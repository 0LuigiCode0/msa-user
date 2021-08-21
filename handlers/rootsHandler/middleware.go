package rootsHandler

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func applyCORS(handler http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(headersOk, originsOk, methodsOk)(handler)
}
