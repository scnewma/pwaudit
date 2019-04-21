SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=pwaudit
PKG="github.com/scnewma/pwaudit"

VERSION="0.0.1"
BUILD_TIME=`date -u +'%Y-%m-%dT%H:%M:%SZ'`

LDFLAGS=-ldflags "-X ${PKG}/pkg/version.Version=${VERSION} -X ${PKG}/pkg/version.BuildDate=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} ./cmd/pwaudit/

install:
	go install ${LDFLAGS} ./cmd/pwaudit/...

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi

.PHONY: install clean
