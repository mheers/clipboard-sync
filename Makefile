## This is a self-documented Makefile. For usage information, run `make help`:
##
## For more information, refer to https://suva.sh/posts/well-documented-makefiles/

SHELL := /bin/bash

all: help

##@ Running
start-server: ## Start the server
	@echo "Starting server..."
	docker-compose up nats

create-seeds:
	@echo "Creating seeds..."
	docker run --rm mheers/nats-seeder seeds

create-credentials:
	nats-seeder \
	user-nkey \
	--operator-seed SOAON2QVZ5L7CMOO5W3PV4F7OCDU7L6AXIO5VA2YWIBTTSLUN64UNOU63M \
	--account-seed SAADBIEN2MTECGRQDK3Y6XHK7PADDSXR6SOCQOM5GFORHLBAX6V6C65OOE \
	-u test \
	-p "\$JS.API.>" -s "\$JS.API.>" -p "_INBOX.>" -s "_INBOX.>" \
	-p "instance.clipboard.*" \
	-s "instance.clipboard.*" \
	> /tmp/test.creds && cat /tmp/test.creds

##@ Building
install-dependencies: ## installs dependencies
	go mod download

set-version: ## Sets the version
	./ci/set-version.sh

build: ## Build the binary
	go build -o clipboardsync .

build-windows:
	GOOS=windows GOARCH=amd64 go build -o clipboardsync.exe .

##@ Testing
test-unit: ## Starts unit tests
	go test ./... -race -coverprofile cover.out
	go tool cover -func cover.out
	rm cover.out

##@ Helpers

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
