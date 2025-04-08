.PHONY: build test lint addlicense install-tools

build:
	CGO_ENABLED=1 go build -v ./...

test:
	CGO_ENABLED=1 go test -race -v ./...

lint:
	go vet ./...
	golangci-lint run ./...
	staticcheck ./...
	gosec -quiet ./...

addlicense:
	addlicense -v -c "voidint <voidint@126.com>" -l mit .

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/google/addlicense@latest