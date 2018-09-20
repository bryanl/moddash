package main

import (
	"encoding/json"

	"github.com/bryanl/moddash/pkg/module"
	"github.com/bryanl/moddash/pkg/proto"
)

func main() {
	module.NewServer("overview", &other{})
}

type other struct{}

func (other) Contents(path string) ([]*proto.Content, error) {
	var contents []*proto.Content

	if path == "" {
		path = "other"
	}

	c := module.Content{
		ContentType: "table",
		Data: map[string]interface{}{
			"columns": []string{"1", "2", "3"},
			"rows": [][]string{
				{"z", "y", "x"},
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

func (other) Metadata() (*proto.Metadata, error) {
	return &proto.Metadata{
		Name:     "Other",
		RootPath: "other",
	}, nil
}

func (other) Navigation() ([]*proto.NavigationEntry, error) {
	entries := []*proto.NavigationEntry{
		{
			Key:  "sub 1",
			Path: "/other/sub-1",
		},
	}
	return entries, nil
}
