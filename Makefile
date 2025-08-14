config ?= configs/local.yaml

build:
	@go build -o bin/wt-guided-weapons cmd/main.go
run: build
	@./bin/wt-guided-weapons -config=$(config)

test:
	@go clean -testcache
	@go test -v -count=1 ./...