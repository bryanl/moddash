package main

import (
	"encoding/json"

	"github.com/bryanl/moddash/pkg/module"
	"github.com/bryanl/moddash/pkg/proto"
)

func main() {
	module.NewServer("overview", &overview{})
}

var (
	contents = map[string][]module.Content{}
)

type overview struct{}

func (overview) Contents(path string) ([]*proto.Content, error) {
	var contents []*proto.Content

	if path == "" {
		path = "overview"
	}

	c := module.Content{
		ContentType: "table",
		Data: map[string]interface{}{
			"columns": []string{"foo", "bar", "baz"},
			"rows": [][]string{
				{"a", "b", "c"},
			},
			"title": path,
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
			Key:  "Workloads",
			Path: "/overview/workloads",
		},
		{
			Key:  "Service",
			Path: "/overview/service",
		},
		{
			Key:  "Config & Storage",
			Path: "/overview/config-and-storage",
		},
		{
			Key:  "Custom Resources",
			Path: "/overview/custom-resources",
		},
		{
			Key:  "RBAC",
			Path: "/overview/rbac",
		},
	}
	return entries, nil
}
