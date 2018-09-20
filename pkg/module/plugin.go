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

func (c *GRPCClient) Contents(path string) ([]*proto.Content, error) {
	req := &proto.ContentRequest{
		Path: path,
	}
	resp, err := c.client.Contents(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp.Contents, nil
}

func (c *GRPCClient) Metadata() (*proto.Metadata, error) {
	resp, err := c.client.Metadata(context.Background(), &proto.Empty{})
	if err != nil {
		return nil, err
	}

	return resp.Metadata, nil
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

func (s *GRPCServer) Contents(ctx context.Context, in *proto.ContentRequest) (*proto.ContentResponse, error) {
	path := in.GetPath()

	contents, err := s.Impl.Contents(path)
	if err != nil {
		return nil, err
	}

	return &proto.ContentResponse{
		Contents: contents,
	}, nil
}

func (s *GRPCServer) Metadata(ctx context.Context, in *proto.Empty) (*proto.MetadataResponse, error) {
	metadata, err := s.Impl.Metadata()
	if err != nil {
		return nil, err
	}

	return &proto.MetadataResponse{
		Metadata: metadata,
	}, nil
}

func (s *GRPCServer) Navigation(ctx context.Context, in *proto.Empty) (*proto.NavigationResponse, error) {
	entries, err := s.Impl.Navigation()

	return &proto.NavigationResponse{
		Entries: entries,
	}, err
}
