package server

import (
	"errors"
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/middleware"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"log"
	"log/slog"
	"net/http"
)

type TibiaMktApiServer struct {
	address string
	router  *http.ServeMux
}

func New(address string) (*TibiaMktApiServer, error) {
	router := http.NewServeMux()

	return &TibiaMktApiServer{
		address: address,
		router:  router,
	}, nil
}

func (s *TibiaMktApiServer) Start() {
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", s.router))

	MuxMiddleWareChain := middleware.NewMiddlewareChain(
		middleware.NewRequestTracingMiddleware,
	)

	server := http.Server{
		Addr:    s.address,
		Handler: MuxMiddleWareChain.ApplyOn(v1),
	}

	slog.Info("Starting the TibiaMktApiServer at " + s.address)

	err := server.ListenAndServe()

	if err != nil {
		evtErr := types.ApiError{Msg: err.Error(), Function: "Start", File: "server.go"}

		log.Fatal(evtErr.Error())
	}
}

func (s *TibiaMktApiServer) Route(url string, handler http.Handler) {
	s.router.Handle(url, handler)
}

func (s *TibiaMktApiServer) NewHandler(handler types.TibiaMktHttpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appError types.ApiErrorInterface

		if err := handler(w, r); errors.As(err, &appError) {
			if appError.IsDomain() {
				response := types.ServerResponse{
					Ok:    false,
					Error: appError.PresentableError(),
				}

				if encodeErr := service.Encode[types.ServerResponse](w, http.StatusOK, response); encodeErr != nil {
					log.Fatal(encodeErr.Error())
				}

				slog.Warn(fmt.Sprintf("%s TraceId=%s", appError.Error(), r.Header.Get("traceId")))

				return
			}

			response := types.ServerResponse{
				Ok:    false,
				Error: constants.ServerGenericError,
			}

			if encodeErr := service.Encode[types.ServerResponse](w, http.StatusInternalServerError, response); encodeErr != nil {
				log.Fatal(encodeErr.Error())
			}

			slog.Error(fmt.Sprintf("%s TraceId=%s", appError.Error(), r.Header.Get("traceId")))
		}
	}
}
