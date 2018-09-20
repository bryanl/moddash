GOCMD=go
GOBUILD=$(GOCMD) build
MODULE_CACHE_DIR=/tmp/module-cache

install-moddash:
	$(GOCMD) install ./cmd/moddash

build-modules: build-module-overview build-module-other

build-module-overview:
	$(GOBUILD) -o $(MODULE_CACHE_DIR)/md-overview ./cmd/md-overview

build-module-other:
	$(GOBUILD) -o $(MODULE_CACHE_DIR)/md-other ./cmd/md-other

generate-proto:
	protoc --go_out=plugins=grpc:. pkg/proto/*.proto

setup-ui: ui-deps run-ui

run-ui:
	cd ui; npm start

ui-deps:
	cd ui; npm i