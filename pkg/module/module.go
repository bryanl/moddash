package module

import (
	"github.com/bryanl/moddash/pkg/proto"
	plugin "github.com/hashicorp/go-plugin"
)

type Module interface {
	Contents(path string) ([]*proto.Content, error)
	Metadata() (*proto.Metadata, error)
	Navigation() ([]*proto.NavigationEntry, error)
}

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "HQ_MODULE",
	MagicCookieValue: "seven",
}
