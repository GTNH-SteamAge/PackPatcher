.DEFAULT_TARGET: all

export version := 1.0.0
export branch := $(shell git rev-parse --abbrev-ref HEAD)
export commit := $(shell git rev-parse --short=8 HEAD)
export internalPKG := github.com/GTNH-SteamAge/PackPatcher/internal
export LDFLAGS := -X $(internalPKG).Version=$(version) -X $(internalPKG).Branch=$(branch) -X $(internalPKG).Commit=$(commit) $(LDFLAGS)

.PHONY: all
all: clean lint build

.PHONY: build
build:
	CGO_ENABLED=0 go build -o packpatcher -ldflags "$(LDFLAGS)"

.PHONY: clean
clean:
	rm -rf packpatcher coverage.out dist/

.PHONY: deps
deps:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.1
	go install github.com/goreleaser/goreleaser/v2@v2.8.0

.PHONY: lint
lint:
	@command -v golangci-lint &>/dev/null || { \
		echo "target requires 'golangci-lint': run make deps"; \
		exit 1; \
	}

	golangci-lint run

.PHONY: release
release: clean
	@command -v goreleaser &>/dev/null || { \
		echo "target requires 'goreleaser': run make deps"; \
		exit 1; \
	}

	goreleaser release

.PHONY: snapshot
snapshot: clean
	@command -v goreleaser &>/dev/null || { \
		echo "target requires 'goreleaser': run make deps"; \
		exit 1; \
	}

	goreleaser release --snapshot
