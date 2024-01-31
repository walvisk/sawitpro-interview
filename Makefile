

.PHONY: clean all init generate generate_mocks

all: build/main

build/main: cmd/main.go generated
	@echo "Building..."
	go build -o $@ $<

clean:
	rm -rf generated

init: generate
	go mod tidy
	go mod vendor

test:
	go test -short -coverprofile coverage.out -v ./...

generate: generated generate_mocks generate_mocks_service

generated: api.yml
	@echo "Generating files..."
	mkdir generated || true
	oapi-codegen --package generated -generate types,server,spec $< > generated/api.gen.go

INTERFACES_GO_FILES := $(shell find repository -name "interfaces.go")
INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:%.go=%.mock.gen.go)

generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))


INTERFACES_SERVICE_GO_FILES := $(shell find service -name "interfaces.go")
INTERFACES_SERVICE_GEN_GO_FILES := $(INTERFACES_SERVICE_GO_FILES:%.go=%.mock.gen.go)

generate_mocks_service: $(INTERFACES_SERVICE_GEN_GO_FILES)
$(INTERFACES_SERVICE_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks service $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))