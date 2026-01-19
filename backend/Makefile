.PHONY: run build docker-build docker-push deps lint lint-fix lint-config test-e2e clean

# Config
IMG ?= cluster-info-backend
TAG ?= latest
LOCALBIN ?= $(shell pwd)/bin

GOLANGCI_LINT_VERSION ?= v2.0.2
GOLANGCI_LINT = $(LOCALBIN)/golangci-lint

GINKGO_VERSION ?= latest
GINKGO = $(LOCALBIN)/ginkgo

# Default target
all: build

# Run the app locally
run:
	go run cmd/main.go

# Build the binary
build: $(LOCALBIN)
	go build -o $(LOCALBIN)/cluster-info-backend cmd/main.go

# Docker image build
docker-build:
	docker build -t $(IMG):$(TAG) .

# Docker image push
docker-push: docker-build
	docker push $(IMG):$(TAG)

# Dependency management
deps:
	go mod tidy
	go mod download

# Create bin dir if needed
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

# Install golangci-lint
$(GOLANGCI_LINT): $(LOCALBIN)
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run

lint-fix: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --fix

lint-config: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) config verify

# Install Ginkgo CLI
$(GINKGO): $(LOCALBIN)
	@echo "Installing Ginkgo CLI..."
	GOBIN=$(LOCALBIN) go install github.com/onsi/ginkgo/v2/ginkgo@$(GINKGO_VERSION)
	GOBIN=$(LOCALBIN) go install github.com/onsi/gomega/...@latest

test-e2e: $(GINKGO)
	$(GINKGO) run -v test/e2e_tests/

# Clean everything
clean:
	rm -rf $(LOCALBIN)/
	docker rmi $(IMG):$(TAG) || true
