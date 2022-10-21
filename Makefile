GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_SRC_PATH=.
GO_BIN_PATH=bin
BINARY_NAME=main
GOPATH:=$(shell go env GOPATH)
GOPACKAGES=$(shell go list ./... | grep -v /proto/ | grep -v /vendor/)
TESTPACKAGES=$(shell go list ./...)
GOFILES= $(shell find . -type f -name '*.go' -not -path "./proto/*" -not -path "./vendor/*")

init:
ifeq ($(OS),Windows_NT)
	@echo "Unsupport OS";
	@exit 1;
else

UNAME_S := $(shell uname -s | tr [A-Z] [a-z])

ifeq ($(UNAME_S),linux)
	OS_NAME := $(shell (cat /etc/*release* | grep "^NAME" | awk -F'"' '{print $$2}' | awk '{print $$1}'))
else ifeq ($(UNAME_S),darwin)
	OS_NAME = $(UNAME_S)
else
	OS_NAME = Unknow
endif

endif

agents: init
	@if ! which protoc &> /dev/null ; then \
		if [ "$(OS_NAME)" = "ubuntu" ] ; then \
			apt-get install protobuf; \
		elif [ "$(OS_NAME)" = "alpine" ] ; then \
			apk add protobuf; \
		elif [ "$(OS_NAME)" = "darwin" ] ; then \
			brew install protobuf; \
		else \
			echo "Unsupport OS"; \
			exit 1; \
		fi; \
	fi;

	GO111MODULE=on ${GO_CMD} install github.com/swaggo/swag/cmd/swag@v1.8.4

.PHONY: run
run:
	godotenv -f local.env go run main.go

.PHONY: load-test
load-test:
	k6 run ./load-tests/add-free.js -q
	k6 run ./load-tests/add-with-redis-lock.js -q

.PHONY: swagger
swagger: agents
	swag init -ot go,yaml -g $(GO_SRC_PATH)/main.go