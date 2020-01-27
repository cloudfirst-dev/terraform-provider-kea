TEST?=github.com/cloudfirst-dev/terraform-provider-kea/kea
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=kea

default: build

build: 
	go install

fmt:
	gofmt -w $(GOFMT_FILES)

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)