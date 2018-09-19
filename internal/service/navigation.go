package service

import (
	"encoding/json"
	"net/http"

	"github.com/bryanl/moddash/pkg/module"
)

type navigation struct {
	loader *module.Loader
}

var _ (http.Handler) = (*navigation)(nil)

func newNavigation(loader *module.Loader) *navigation {
	return &navigation{
		loader: loader,
	}
}

func (n *navigation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	entries, err := n.loader.NavigationEntries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	json.NewEncoder(w).Encode(entries)
}
