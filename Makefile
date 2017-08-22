BINARY=mongosnap
VERSION=0.0.2
BUILD=`git rev-parse HEAD`
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
	ghr -t ${GITHUB_TOKEN} -u ${CIRCLE_PROJECT_USERNAME} -r ${CIRCLE_PROJECT_REPONAME} --replace ${BUILD} dist/
