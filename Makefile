VERSION := $(shell git describe --tags --always 2>/dev/null | sed 's/-g[0-9a-f]*$$//' || echo "0.1.0-dev")

RUNNER ?=

BUILD_DIR ?= $(if $(GITHUB_WORKSPACE),$(GITHUB_WORKSPACE)/build,/tmp/wardrobe-build-dir)
export BUILD_DIR

BINARY = $(BUILD_DIR)/wardrobe

PACKAGES = *.deb *.rpm *.pkg.tar.zst PKGBUILD

# -----------------------------------------------------------
all: build docs
	@echo "--------------------------------------"
	@echo "Wardrobe built successfully! 🐧👗"
	@echo "Version: $(VERSION)"
	@echo "Binary:  $(BINARY)"
	@echo "--------------------------------------"

# -----------------------------------------------------------
# Build
# -----------------------------------------------------------
build: | $(BUILD_DIR)
	@echo "  MAKING wardrobe (Go)..."
	@go build -ldflags "-X 'github.com/pieroproietti/penguins-wardrobe/cmd.Version=$(VERSION)'" -o $(BINARY) main.go

$(BUILD_DIR):
	@mkdir -p $@

# -----------------------------------------------------------
# Docs & packaging
# -----------------------------------------------------------
docs: build
	@echo "  GENERATING DOCUMENTATION & COMPLETIONS..."
	@mkdir -p docs
	@-$(RUNNER) $(BINARY) _gen_docs --target $(BUILD_DIR)/docs

package: all
	@echo "  PACKAGING NATIVE OS DISTRIBUTION..."
	@BUILD_DIR=$(BUILD_DIR) PROJ_ROOT=$(PWD) $(RUNNER) $(BINARY) tools build

# -----------------------------------------------------------
# Clean
# -----------------------------------------------------------
clean:
	@echo "  Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR) wardrobe
	@rm -f $(PACKAGES)
	@rm -rf docs/man docs/completion docs/md

install: build
	sudo install -m 755 $(BINARY) /usr/local/bin/wardrobe

test:
	go test -v ./...

.PHONY: all build docs package clean install test
