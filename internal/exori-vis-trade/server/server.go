package server

import (
	"github.com/adriein/exori-vis-trade/pkg/types"
	"log"
	"log/slog"
	"net/http"
)

type ExoriVisTradeApiServer struct {
	address string
	router  *http.ServeMux
}

func New(address string) (*ExoriVisTradeApiServer, error) {
	router := http.NewServeMux()

	return &ExoriVisTradeApiServer{
		address: address,
		router:  router,
	}, nil
}

func (s *ExoriVisTradeApiServer) Start() {
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", s.router))

	server := http.Server{
		Addr:    s.address,
		Handler: v1,
	}

	slog.Info("Starting the ExoriVisTradeApiServer at " + s.address)

	err := server.ListenAndServe()

	if err != nil {
		evtErr := &types.EvtError{Msg: err.Error(), Function: "Start", File: "server.go"}

		log.Fatal(evtErr.Error())
	}
}

func (s *ExoriVisTradeApiServer) Route(url string, handler http.HandlerFunc) *ExoriVisTradeApiServer {
	s.router.HandleFunc(url, handler)

	return s
}
