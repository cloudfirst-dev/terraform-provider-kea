TEST?=github.com/cloudfirst-dev/terraform-provider-kea/kea
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=kea

default: build

build: 
	go install

fmt:
	gofmt -w $(GOFMT_FILES)

test-compile:
	go test ./...

deploy:
	make build
	cp $$GOPATH/bin/terraform-provider-kea ~/.terraform.d/plugins/