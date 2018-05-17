SDK_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/" | grep -v "/example")
SDK_TEST_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/" | grep -v "/config" | grep -v "/example")

all: build unit

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  build                   to go build the SDK"
	@echo "  unit                    to run unit tests"
	@echo "  lint                    to lint the SDK"
	@echo "  generate                to generate the Go Structs from JSON schema"

build:
	@echo "go build SDK and vendor packages"
	@go build ${SDK_ONLY_PKGS}

unit:  build 
	@echo "go test SDK  package"
	@go test  -v $(SDK_TEST_ONLY_PKGS)

lint:  build 
	@echo "go lint netscaler package (ignoring generated packages)"
	@golint   netscaler | grep -v netscaler/resources.go || true

generate:
	@echo "Generate go schema from json schema"
	(cd tools; ./generate.sh)

