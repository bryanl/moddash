package service

import (
	"context"
	"log"
	"net/http"

	"github.com/bryanl/moddash/pkg/module"
)

type Service struct {
	loader *module.Loader
}

func New(loader *module.Loader) *Service {
	return &Service{
		loader: loader,
	}
}

func (s *Service) Run(ctx context.Context) {
	m := http.NewServeMux()
	m.Handle("/navigation", newNavigation(s.loader))

	server := http.Server{Addr: ":8000", Handler: m}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	select {
	case <-ctx.Done():
		server.Shutdown(ctx)
	}
}
