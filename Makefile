SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=gherkinize
VERSION=0.0.1
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-X github.com/wesovilabs/gherkinize/core.Version=${VERSION} -X github.com/wesovilabs/gherkinize/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
		go build ${LDFLAGS} -o dist/${BINARY} gherkinize.go

.PHONY: install
install:
		go install ${LDFLAGS} ./...

.PHONY: vet
vet:
		go vet

.PHONY: clean
clean:
		if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: test
test:
		go test ${LDFLAGS} ./...

.PHONY: get-deps
get-deps:
		go get github.com/fatih/color
		go get github.com/urfave/cli
		go get github.com/BurntSushi/toml
