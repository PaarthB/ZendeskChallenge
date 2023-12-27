# TOP gives us the root of the repository
TOP  = $(shell git rev-parse --show-toplevel)

# BIN is the location of generated files
BIN  = $(TOP)/bin
SHELL := bash

# DIST is the location of files for distribution
DIST = $(BIN)/dist

GOARCH      ?= $(shell go env GOARCH)
GOOS        ?= $(shell go env GOOS)
GOPATH      ?= $(shell go env GOPATH)
GOEXE       ?= $(shell GOOS=$(GOOS) go env GOEXE)
GOCMD   	?=  go
GOBUILD = $(GOCMD) build


###################################################################################################

VERSION     ?= $(shell cat .version.txt)
BUILD_DATE  = $(shell date +%Y%m%d.%H%M)
PACKAGE    = $(shell go list)
NAME       = $(notdir $(PACKAGE))
EXEC       = cli$(GOEXE)
BUNDLE     = $(NAME)-$(VERSION)-$(GOOS)-$(GOARCH).tar.gz
BUNDLE_SHA = $(BUNDLE).sha256

BINDIR     = $(TOP)/bin/$(GOOS)-$(GOARCH)

PLUGIN_PLATFORMS    = darwin-amd64 linux-amd64 windows-amd64 darwin-arm64 linux-arm64
PLUGIN_DIST         = $(DIST)/$(NAME)
PLUGIN_RELEASE      = $(PLUGIN_DIST)/$(VERSION)

PLUGIN_EXECUTABLE   = $(TOP)/bin/$(GOOS)-$(GOARCH)/$(EXEC)

PLUGIN_BUNDLE       = $(PLUGIN_RELEASE)/$(GOOS)-$(GOARCH).tar.gz
PLUGIN_BUNDLE_SHA   = $(PLUGIN_BUNDLE).sha256

PLUGIN_BUNDLES      = $(addsuffix .tar.gz, $(PLUGIN_PLATFORMS))
PLUGIN_MANIFESTS    = $(addsuffix .sha256, $(PLUGIN_BUNDLES))

PLUGIN_RELEASE_BUNDLE = $(PLUGIN_RELEASE).tar.gz

###################################################################################################

$(BIN):
	mkdir -p $@

$(DIST):
	mkdir -p $@

###################################################################################################

$(PLUGIN_RELEASE):
	mkdir -p $@

$(BINDIR):
	mkdir -p $@

.PHONY: $(BINDIR)/$(EXEC)
$(BINDIR)/$(EXEC): $(BINDIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -ldflags "-X main.Version=$(VERSION)" -o $@ $(PACKAGE)

.DEFAULT_GOAL := default

.PHONY: default
default: build

.PHONY: build
build: $(BINDIR)/$(EXEC)


.PHONY: clean
clean:
	-rm -f $(BINDIR)/$(EXEC)

###################################################################################################

.PHONY: test
test:
	TEST_ENV=true GOOS=$(GOOS) GOARCH=$(GOARCH) go test -coverpkg= ./... -v


###################################################################################################

$(BUNDLE): build
	tar -C $(BINDIR) -czf $(BUNDLE) $(EXEC)

$(BUNDLE_SHA): $(BUNDLE)
	openssl sha256 -hex -r < $(BUNDLE) | cut -f 1 -d " " > $(BUNDLE_SHA)

.PHONY: bundle
bundle: $(BUNDLE) $(BUNDLE_SHA)
###################################################################################################

$(PLUGIN_BUNDLE): $(PLUGIN_RELEASE) $(PLUGIN_EXECUTABLE)
	tar -C $(dir $(PLUGIN_EXECUTABLE)) -czf $@ $(EXEC)

$(PLUGIN_BUNDLE_SHA): $(PLUGIN_BUNDLE)
	openssl sha256 -hex -r < $(PLUGIN_BUNDLE) | cut -f 1 -d " " > $(PLUGIN_BUNDLE_SHA)

###################################################################################################

.PHONY: wrapper
wrapper:
	make -C wrapper wrapper

###################################################################################################

.PHONY: $(PLUGIN_BUNDLES)

darwin-amd64.tar.gz: GOOS=darwin
darwin-amd64.tar.gz: GOARCH=amd64
darwin-amd64.tar.gz:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(MAKE) bundle

linux-amd64.tar.gz: GOOS=linux
linux-amd64.tar.gz: GOARCH=amd64
linux-amd64.tar.gz:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(MAKE) bundle

windows-amd64.tar.gz: GOOS=windows
windows-amd64.tar.gz: GOARCH=amd64
windows-amd64.tar.gz:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(MAKE) bundle

darwin-arm64.tar.gz: GOOS=darwin
darwin-arm64.tar.gz: GOARCH=arm64
darwin-arm64.tar.gz:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(MAKE) bundle

linux-arm64.tar.gz: GOOS=linux
linux-arm64.tar.gz: GOARCH=arm64
linux-arm64.tar.gz:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(MAKE) bundle

###################################################################################################

.PHONY: bundle bundle-all bundle-release

bundle: $(PLUGIN_BUNDLE) $(PLUGIN_BUNDLE_SHA)

bundle-all: $(PLUGIN_BUNDLES)

$(PLUGIN_RELEASE_BUNDLE): bundle-all
	tar -C $(dir $(PLUGIN_RELEASE)) -czf $@ $(addprefix $(VERSION)/,$(PLUGIN_BUNDLES)) $(addprefix $(VERSION)/,$(PLUGIN_MANIFESTS))

bundle-release: $(PLUGIN_RELEASE_BUNDLE)

###################################################################################################

.PHONY: release

release: bundle-release

setup:
	cp $(BINDIR)/$(EXEC) .