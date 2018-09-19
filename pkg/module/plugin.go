package module

import (
	"context"

	"github.com/bryanl/moddash/pkg/proto"
	plugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type Plugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Module
}

var _ plugin.GRPCPlugin = (*Plugin)(nil)

func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterModuleServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewModuleClient(c)}, nil
}

type GRPCClient struct {
	client proto.ModuleClient
}

func (c *GRPCClient) Navigation() ([]*proto.NavigationEntry, error) {
	resp, err := c.client.Navigation(context.Background(), &proto.Empty{})
	if err != nil {
		return nil, err
	}

	return resp.Entries, nil
}

type GRPCServer struct {
	Impl Module
}

var _ proto.ModuleServer = (*GRPCServer)(nil)

func (s *GRPCServer) Navigation(ctx context.Context, in *proto.Empty) (*proto.NavigationResponse, error) {
	entries, err := s.Impl.Navigation()

	return &proto.NavigationResponse{
		Entries: entries,
	}, err
}
