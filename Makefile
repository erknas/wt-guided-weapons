build:
	@go build -o bin/wt-guided-weapons cmd/main.go
run: build
	@./bin/wt-guided-weapons