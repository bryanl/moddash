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
			"other": &module.Plugin{Impl: &overview{}},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}

type overview struct{}

func (overview) Metadata() (*proto.Metadata, error) {
	return &proto.Metadata{
		Name:     "Overview",
		RootPath: "/overview",
	}, nil
}

func (overview) Navigation() ([]*proto.NavigationEntry, error) {
	entries := []*proto.NavigationEntry{
		{
			Name: "Workloads",
			Path: "/workloads",
		},
		{
			Name: "Service",
			Path: "/service",
		},
		{
			Name: "Config & Storage",
			Path: "/config",
		},
		{
			Name: "Custom Resources",
			Path: "/custom",
		},
		{
			Name: "RBAC",
			Path: "/rbac",
		},
	}
	return entries, nil
}
