BINARY_NAME=webhook-app

ifeq ($(OS),Windows_NT)
    SHELL := pwsh.exe
else
   SHELL := pwsh
endif
.SHELLFLAGS := -NoProfile -Command

build:
	go build -o build/${BINARY_NAME}-linux64 ./api/cmd/main.go
	go build -o build/${BINARY_NAME}-win64.exe ./api/cmd/main.go

run-linux:
	$(CURDIR)/${BINARY_NAME}-linux64

run-windows:
	$(CURDIR)/${BINARY_NAME}-win64.exe

run-dev:generate-templ-docker
	go run ./api/cmd/main.go

generate-templ-docker:
	docker run -v  $(CURDIR):/app -w=/app ghcr.io/a-h/templ:latest generate

clean:
	go clean
	rm "${BINARY_NAME}-*"

deps:
	go mod download

vet:
	go vet