SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY := pwaudit
PKG := github.com/scnewma/pwaudit

VERSION = "0.0.1"
BUILD_TIME = `date -u +'%Y-%m-%dT%H:%M:%SZ'`

LDFLAGS=-ldflags "-X ${PKG}/pkg/version.Version=${VERSION} -X ${PKG}/pkg/version.BuildDate=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build ${LDFLAGS} -o bin/${BINARY} ./cmd/pwaudit/

install:
	go install ${LDFLAGS} ./cmd/pwaudit/...

clean:
	rm -rf bin

PLATFORMS := windows linux darwin
os = $(word 1, $@)

$(PLATFORMS):
	mkdir -p bin
	GOOS=$(os) GOARCH=amd64 go build ${LDFLAGS} -o bin/$(BINARY)-v$(VERSION)-$(os)-amd64 ./cmd/pwaudit/

release: windows linux darwin

.PHONY: install clean release windows linux darwin
