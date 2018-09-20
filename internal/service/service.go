package service

import (
	"context"
	"log"
	"net/http"

	"github.com/bryanl/moddash/pkg/module"
)

// Service is the cli tool's API service.
type Service struct {
	loader *module.Loader
}

// New creates an instance of Service.
func New(loader *module.Loader) *Service {
	return &Service{
		loader: loader,
	}
}

// Run runs the service.
func (s *Service) Run(ctx context.Context) {
	m := http.NewServeMux()
	m.Handle("/navigation", newNavigation(s.loader))
	m.Handle("/contents/", newContents(s.loader))

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
