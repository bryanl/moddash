package module

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/pkg/errors"
)

func AvailableModules(moduleCacheDir string) ([]string, error) {
	fis, err := ioutil.ReadDir(moduleCacheDir)
	if err != nil {
		return nil, err
	}

	var list []string

	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}

		if fi.Mode()&0111 == 0 {
			continue
		}

		list = append(list, filepath.Join(moduleCacheDir, fi.Name()))

	}

	return list, nil
}

type Client struct {
	Name   string
	Module Module

	client *plugin.Client
}

func NewClient(modulePath string) (*Client, error) {
	_, file := filepath.Split(modulePath)

	pluginMap := map[string]plugin.Plugin{
		file: &Plugin{},
	}

	pluginClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  Handshake,
		Plugins:          pluginMap,
		Cmd:              exec.Command(modulePath),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})

	rpcClient, err := pluginClient.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcClient.Dispense(file)
	if err != nil {
		return nil, err
	}

	m, ok := raw.(Module)
	if !ok {
		return nil, errors.New("module not found")
	}

	return &Client{
		Name:   file,
		Module: m,

		client: pluginClient,
	}, nil
}

func (c *Client) Close() {
	c.client.Kill()
}
