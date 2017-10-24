TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build

.PHONY: build
build: fmtcheck ## Build the provider
	go install

.PHONY: test
test: fmtcheck ## Test the provider
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

.PHONY: testacc
testacc: fmtcheck ## Test acceptance of the provider
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

.PHONY: vet
vet: ## Run go vet across the provider
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: fmt
fmt: ## Run gofmt across all go files
	gofmt -w $(GOFMT_FILES)

.PHONY: fmtcheck
fmtcheck: ## Check that code complies with gofmt requirements
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: errcheck
errcheck: ## Check for unchecked errors
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

.PHONY: test-compile
test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./aws"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: help
help:
	@echo "Valid targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
