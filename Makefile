SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=gherkinize
VERSION=0.0.1
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-X github.com/wesovilabs/gherkinize/core.Version=${VERSION} -X github.com/wesovilabs/gherkinize/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
		env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o dist/${BINARY} gherkinize.go

.PHONY: linux-arm
linux-arm:
		env GOOS=linux GOARCH=arm go build ${LDFLAGS} -o dist/linux/arm/${BINARY} gherkinize.go

.PHONY: linux-386
linux-386:
		env GOOS=linux GOARCH=386 go build ${LDFLAGS} -o dist/linux/arm/${BINARY} gherkinize.go

.PHONY: linux-amd64
linux-amd64:
		env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o dist/linux/arm/${BINARY} gherkinize.go


.PHONY: darwin-arm
darwin-arm:
		env GOOS=darwin GOARCH=arm go build ${LDFLAGS} -o dist/darwin/arm/${BINARY} gherkinize.go

.PHONY: darwin-386
darwin-386:
		env GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o dist/darwin/arm/${BINARY} gherkinize.go

.PHONY: darwin-amd64
darwin-amd64:
		env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o dist/darwin/arm/${BINARY} gherkinize.go

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
