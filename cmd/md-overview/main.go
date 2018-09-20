package main

import (
	"encoding/json"

	"github.com/bryanl/moddash/pkg/module"
	"github.com/bryanl/moddash/pkg/proto"
)

func main() {
	module.NewServer("overview", &overview{})
}

type overview struct{}

func (overview) Contents(path string) ([]*proto.Content, error) {
	var contents []*proto.Content

	c := module.Content{
		ContentType: "table",
		Data: map[string]interface{}{
			"columns": []string{"foo", "bar", "baz"},
			"rows": [][]string{
				{"a", "b", "c"},
			},
			"title": "table 1",
		},
	}

	data, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	content := &proto.Content{
		Data: data,
	}

	contents = append(contents, content)

	return contents, nil
}

func (overview) Metadata() (*proto.Metadata, error) {
	return &proto.Metadata{
		Name:     "Overview",
		RootPath: "overview",
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
