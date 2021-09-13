LDFLAGS := -s -w

build:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ./build/enorith ./cmd/app
	docker build -t enorith .