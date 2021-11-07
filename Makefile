.PHONY: build release 
GIT_HASH := $(shell git rev-parse HEAD)
RELEASE_VERSION := ${RELEASE_VERSION}

GIT_DIRTY := yes
ifeq ($(shell git status --porcelain),"")
    GIT_DIRTY = no
endif


build:
	CGO_ENABLED=0 go build \
		-ldflags "-X main.gitHash=${GIT_HASH} -X main.releaseVersion=${RELEASE_VERSION} -X main.gitDirty=${GIT_DIRTY}" \
		$(BUILD_FLAGS)

release:
	mkdir -p build
	GOOS=linux GOARCH=amd64  BUILD_FLAGS="-o ./build/go-ini-query-linux-amd64" make build
	GOOS=linux GOARCH=arm64  BUILD_FLAGS="-o ./build/go-ini-query-linux-arm64" make build
	GOOS=linux GOARCH=386    BUILD_FLAGS="-o ./build/go-ini-query-linux-386" make build
	GOOS=darwin GOARCH=amd64 BUILD_FLAGS="-o ./build/go-ini-query-darwin-amd64" make build
	GOOS=darwin GOARCH=arm64 BUILD_FLAGS="-o ./build/go-ini-query-darwin-arm64" make build