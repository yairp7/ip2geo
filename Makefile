include .env

GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BUILD_DIR=build
BINARY_NAME=server

run:
	$(GOCMD) run .

build:
	mkdir -p $(BUILD_DIR)/
	$(GOCMD) build -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	rm -fr ./$(BUILD_DIR)

watch:
	$(eval PACKAGE_NAME=$(shell head -n 1 go.mod | cut -d ' ' -f2))
	docker run -it --rm -w /go/src/$(PACKAGE_NAME) -v $(shell pwd):/go/src/$(PACKAGE_NAME) -p ${PORT}:${PORT} cosmtrek/air

test: 
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

docker-build:
	docker build --rm -t $(BINARY_NAME) .

docker-run:
	docker run -p ${PORT}:${PORT} -t $(BINARY_NAME)