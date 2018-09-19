package main

import (
	"github.com/bryanl/moddash/pkg/module"
	"github.com/bryanl/moddash/pkg/proto"
	plugin "github.com/hashicorp/go-plugin"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: module.Handshake,
		Plugins: map[string]plugin.Plugin{
			"overview": &module.Plugin{Impl: &overview{}},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}

type overview struct{}

func (overview) Navigation() ([]*proto.NavigationEntry, error) {
	entries := []*proto.NavigationEntry{
		{
			Name: "Overview",
			Path: "/overview",
		},
	}
	return entries, nil
}
