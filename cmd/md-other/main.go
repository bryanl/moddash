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
			"other": &module.Plugin{Impl: &other{}},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}

type other struct{}

func (other) Metadata() (*proto.Metadata, error) {
	return &proto.Metadata{
		Name:     "Other",
		RootPath: "/other",
	}, nil
}

func (other) Navigation() ([]*proto.NavigationEntry, error) {
	entries := []*proto.NavigationEntry{
		{
			Name: "sub 1",
			Path: "/other",
		},
	}
	return entries, nil
}
