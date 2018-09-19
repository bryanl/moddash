GOCMD=go
GOBUILD=$(GOCMD) build
MODULE_CACHE_DIR=/tmp/module-cache

build-module-overview:
	$(GOBUILD) -o $(MODULE_CACHE_DIR)/md-overview ./cmd/md-overview
