package main

import (
	"fmt"
	"os"

	"github.com/bryanl/moddash/pkg/module"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	moduleList, err := module.AvailableModules("/tmp/module-cache")
	if err != nil {
		os.Exit(1)
	}

	for _, path := range moduleList {
		c, err := module.NewClient(path)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}

		defer c.Close()

		entries, err := c.Module.Navigation()
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}

		spew.Dump(entries)
	}
}
