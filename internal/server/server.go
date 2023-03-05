package server

import (
	"net/http"

	"github.com/syntro-org/newsdeck/config"
	"github.com/syntro-org/newsdeck/internal/handler"
)

type Server struct {
	engine *http.Server
}

func (s *Server) Run(address string) error {
	s.engine = &http.Server{
		Addr:    address,
		Handler: handler.Init(),
	}

	return s.engine.ListenAndServe()
}

func MakeAddr() string {
	conf := config.New()

	if conf.AppMode == config.RELEASE_MODE {
		return ":" + conf.Port
	}

	return "localhost:" + conf.Port
}
