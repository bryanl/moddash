package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bryanl/moddash/internal/service"
	"github.com/bryanl/moddash/pkg/module"
)

func main() {
	var moduleCachePath string
	flag.StringVar(&moduleCachePath, "moduleCachePath", "/tmp/module-cache", "module cache path")

	flag.Parse()

	if err := run(moduleCachePath); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}

func run(moduleCachePath string) error {
	loader, err := module.NewLoader(moduleCachePath)
	if err != nil {
		return err
	}

	defer loader.Close()

	ctx := context.Background()

	go func(ctx context.Context, loader *module.Loader) {
		s := service.New(loader)
		s.Run(ctx)
	}(ctx, loader)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	<-sigCh

	log.Println("waiting for modules to close")
	time.Sleep(3 * time.Second)

	return nil
}
