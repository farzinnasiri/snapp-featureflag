# constants
APP_NAME := featureflag
APP_VERSION := 0.1.0
ROOT := github.com/$(OWNER)/$(APP_NAME)
BINARY_NAME=featureflag
BINARY_PATH=bin

run: ## to start the application
	@go run ./cmd/$(APP_NAME)

wire: | .pre-check-wire ## to generate wire file for dependency injection
	@wire ./internal/app/$(APP_NAME)

.pre-check-wire:
	@if [ -z "$$(which wire)" ]; then \
  		go get -v github.com/google/wire/cmd/wire; \
  	fi

dependencies: ## to install the dependencies
	@go mod tidy -compat=1.18
	@go mod download
	@go mod vendor

clean:
	@rm -rf $(APP_NAME)
	@rm -rf $(BINARY_PATH) || true

.pre-check-formatting-tools:
	@if [ -z "$$(which gofumpt)" ]; then go install mvdan.cc/gofumpt@latest; fi
.PHONY: .pre-check-formatting-tools

fmt: | .pre-check-formatting-tools ## to run go formatter on all source codes across the project
	@gofumpt -l -w './'
.PHONY: fmt

build:
	@echo "Cleanning..."
	@mkdir $(BINARY_PATH)
	@echo "Building featureflag..."
	@CGO_ENABLED=0 go build  -o $(BINARY_PATH)/$(APP_NAME) -ldflags '-w' -v ./cmd/$(APP_NAME)
	@echo "Building is finished"
