.PHONY:frontend
GO           ?= go
GOFMT        ?= $(GO)fmt
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))

pkgs          = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX                  ?= $(shell pwd)
DIRNAME                 ?= $(shell dirname $(shell pwd))

#TAG                     ?= $(shell date +%s)
TAG                     ?= $(shell git rev-parse --short HEAD)
ENV 					?= dev

RUN_ENV                 ?= test

style:
	@echo ">> checking code style"
	@! $(GOFMT) -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

.PHONY: build-cmd
build-cmd:
	${MAKE} -C cmd build

build: build-cmd

image: build-cmd
	@echo ">> start building fiber-demo image ..."
	@docker build -t zbd20/fiber-demo:${VERSION} .
	@echo ">> docker image has been built."
	@echo ">> push image to the repository ..."
	@docker push zbd20/fiber-demo:${VERSION}
	@echo ">> completed."

clean:
	@echo ">> remove fiber-demo"
	@rm fiber-demo

swagger:
	@echo ">> swag init"
	@swag init --dir cmd,internal --output api/swagger
