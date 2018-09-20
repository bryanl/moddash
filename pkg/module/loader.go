package module

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/bryanl/moddash/pkg/proto"
	"github.com/pkg/errors"
)

type Loader struct {
	clients []Client
}

func NewLoader(moduleCachePath string) (*Loader, error) {
	moduleList, err := AvailableModules(moduleCachePath)
	if err != nil {
		return nil, err
	}

	var clients []Client

	for _, path := range moduleList {
		c, err := NewClient(path)
		if err != nil {
			return nil, err
		}

		clients = append(clients, *c)
	}

	loader := &Loader{
		clients: clients,
	}

	return loader, nil
}

func (l *Loader) Contents(path string) ([]Content, error) {
	parts := strings.Split(path, "/")

	// find client
	for _, client := range l.clients {
		if client.Name != parts[0] {
			continue
		}

		dataContents, err := client.Module.Contents(strings.Join(parts[1:], "/"))
		if err != nil {
			return nil, err
		}

		var contents []Content
		for _, dataContent := range dataContents {
			var c Content
			if err = json.Unmarshal(dataContent.Data, &c); err != nil {
				return nil, err
			}

			contents = append(contents, c)
		}

		return contents, nil
	}

	return nil, errors.New("not found")
}

func (l *Loader) NavigationEntries() ([]*proto.NavigationEntry, error) {
	var entries []*proto.NavigationEntry

	for _, client := range l.clients {
		metadata, err := client.Module.Metadata()
		if err != nil {
			return nil, err
		}

		subEntries, err := client.Module.Navigation()
		if err != nil {
			return nil, err
		}

		entry := &proto.NavigationEntry{
			Name: metadata.Name,
			Path: client.Name,
			Subs: subEntries,
		}

		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Name == "Overview" {
			return true
		}

		return entries[i].Name < entries[j].Name
	})

	return entries, nil
}

func (l *Loader) Close() {
	for _, client := range l.clients {
		client.Close()
	}
}
