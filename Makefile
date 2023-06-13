buildAt=`date +%Y/%m/%d-%H:%M:%S`
gitHash=`git rev-parse HEAD`
version=`git rev-parse --abbrev-ref HEAD | grep -v HEAD || git describe --exact-match HEAD || git rev-parse HEAD`  ## todo: use current release git tag
flags="-X 'atom/http/utils.Version=${version}' -X 'atom/http/utils.BuildAt=${buildAt}' -X 'atom/http/utils.GitHash=${gitHash}'"
release_flags="-w -s ${flags}"

GOPATH:=$(shell go env GOPATH)

.PHONY: doc
doc:
	@swag fmt
	@swag init -ot json

.PHONY: generate
generate:
	@go generate ./...

.PHONY: tidy
tidy: doc generate
	@go mod tidy

.PHONY: dist
dist:doc generate
	@echo "building..."
	@go build -ldflags=${flags} -o bin/debug/atom
	@cp config.toml bin/debug/

.PHONY: run
run: doc generate
	@echo "building..."
	@go run .

.PHONY: release
release: doc generate
	@go build -ldflags=${flags} -o bin/release/atom
	@cp config.toml bin/release/

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: proto
proto:
	@buf generate

.PHONY: mup
mup:
	@go run . migrate up

.PHONY: mdown
mdown:
	@go run . migrate down

.PHONY: model
model:
	@go run . model
