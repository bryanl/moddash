GOCMD=go
GOBUILD=$(GOCMD) build
MODULE_CACHE_DIR=/tmp/module-cache

build-modules: build-module-overview build-module-other

build-module-overview:
	$(GOBUILD) -o $(MODULE_CACHE_DIR)/md-overview ./cmd/md-overview

build-module-other:
	$(GOBUILD) -o $(MODULE_CACHE_DIR)/md-other ./cmd/md-other