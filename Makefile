BINARY=mongosnap
VERSION=0.0.2
BUILD=`git rev-parse HEAD`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT_GOAL: build

build:
	gox ${LDFLAGS} -o dist/${BINARY}_{{.OS}}_{{.Arch}}"

install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY}; fi

.PHONY: clean install

test:
	go test -v ./...
