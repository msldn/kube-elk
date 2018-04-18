SHELL="/bin/bash"
PLATFORM=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

NAME=elker
VERSION=1.0
DEV_RUN_OPTS ?= consul:

default: build validate test

get-deps:
	go get -u github.com/golang/lint/golint honnef.co/go/tools/cmd/megacheck

build:
	go fmt ./...
	mkdir -p build
	docker build -t $(NAME):$(VERSION) .
	docker save $(NAME):$(VERSION) | gzip -9 > build/$(NAME)_$(VERSION).tgz

licenseok:
	go build ./hack/licenseok

validate: build licenseok
	./hack/lint.bash
	./hack/validate-vendor.bash
	./hack/validate-licence.bash

test:
	./hack/test.bash

install: build
	cp ./dep $(GOBIN)

docusaurus:
	docker run --rm -it -v `pwd`:/dep -p 3000:3000 \
		-w /dep/website node \
		bash -c "npm i --only=dev && npm start"

.PHONY: build validate test install docusaurus
