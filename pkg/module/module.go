package module

import (
	"github.com/bryanl/moddash/pkg/proto"
	plugin "github.com/hashicorp/go-plugin"
)

type Module interface {
	Metadata() (*proto.Metadata, error)
	Navigation() ([]*proto.NavigationEntry, error)
}

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "handshake",
}
