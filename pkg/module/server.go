package module

import (
	plugin "github.com/hashicorp/go-plugin"
)

// NewServer creates a module server.
func NewServer(name string, impl Module) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			name: &Plugin{Impl: impl},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}
