package middleware

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"log"
	"net/http"
	"os"
)

func NewAuthMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerApiKey := r.Header.Get(constants.ApiKeyHeader)
		validApiKey := os.Getenv(constants.TibiaMktApiKey)

		if headerApiKey != validApiKey {
			response := types.ServerResponse{
				Ok: false,
			}

			if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
				log.Fatal(encodeErr.Error())
			}

			return
		}

		handler.ServeHTTP(w, r)
	})
}
