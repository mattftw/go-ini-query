.PHONY: build release test
GIT_HASH := $(shell git rev-parse HEAD)
RELEASE_VERSION := ${RELEASE_VERSION}
GIT_STATE=$(shell (git status --porcelain | grep -q .) && echo dirty || echo clean)

test:
	mkdir -p ./test-output
	go test -coverprofile ./test-output/coverprofile
	go tool cover -html=./test-output/coverprofile -o ./test-output/coverage.html;

build:
	CGO_ENABLED=0 go build \
		-ldflags "-X main.gitHash=${GIT_HASH} -X main.releaseVersion=${RELEASE_VERSION} -X main.gitState=${GIT_STATE}" \
		$(BUILD_FLAGS)

release:
ifeq ($(GIT_STATE),dirty)
	@echo Attempting to build release with dirty repo. Failing.
	@exit 1
endif
	mkdir -p build
	GOOS=linux  GOARCH=amd64  BUILD_FLAGS="-o ./build/go-ini-query-linux-amd64"  make build
	GOOS=linux  GOARCH=arm64  BUILD_FLAGS="-o ./build/go-ini-query-linux-arm64"  make build
	GOOS=linux  GOARCH=386    BUILD_FLAGS="-o ./build/go-ini-query-linux-386"    make build
	GOOS=darwin GOARCH=amd64  BUILD_FLAGS="-o ./build/go-ini-query-darwin-amd64" make build
	GOOS=darwin GOARCH=arm64  BUILD_FLAGS="-o ./build/go-ini-query-darwin-arm64" make build
