BINARY=mongosnap
VERSION=0.0.1
BUILD=`git rev-parse HEAD`
BRANCH=$(shell git symbolic-ref --short HEAD)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT_GOAL: build

build:
	gox ${LDFLAGS} -osarch="linux/amd64" -output "dist/${BINARY}_{{.OS}}_{{.Arch}}"

install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY}; fi

.PHONY: clean install

test:
	go test -v ./...

publish:
ifeq ($(BRANCH),master)
	@echo On branch master, publishing release
	ghr \
	-t ${GITHUB_TOKEN} \
	-u ${CIRCLE_PROJECT_USERNAME} \
	-r ${CIRCLE_PROJECT_REPONAME} \
	--replace \
	v${VERSION} \
	dist/
else
	@echo On branch $(BRANCH), publishing prerelease
	ghr \
	-t ${GITHUB_TOKEN} \
	-u ${CIRCLE_PROJECT_USERNAME} \
	-r ${CIRCLE_PROJECT_REPONAME} \
	--replace \
	--prerelease \
	v${VERSION} \
	dist/
endif

versioncheck:
	@echo ${VERSION}
