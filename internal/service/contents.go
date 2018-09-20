package service

import (
	"encoding/json"
	"net/http"

	"github.com/bryanl/moddash/pkg/module"
)

type contents struct {
	loader *module.Loader
}

var _ (http.Handler) = (*contents)(nil)

func newContents(loader *module.Loader) *contents {
	return &contents{
		loader: loader,
	}
}

func (n *contents) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/contents/"):]
	contents, err := n.loader.Contents(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	json.NewEncoder(w).Encode(contents)
}
