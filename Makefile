SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=gherkinize
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X github.com/wesovilabs/gherkinize/core.Version=${VERSION} -X github.com/wesovilabs/gherkinize/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
		go build ${LDFLAGS} -o ${BINARY} gherkinize.go

.PHONY: install
install:
		go install ${LDFLAGS} ./...

.PHONY: clean
clean:
		if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi