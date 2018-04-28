BINARY = eva
VET_REPORT = vet.report
TEST_REPORT = tests.xml
GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

GOPATH?=${HOME}/go

# Symlink into GOPATH
GITHUB_USERNAME=iopipe
BUILD_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${BINARY}
CURRENT_DIR=$(shell pwd)
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: link clean goget test vet linux darwin windows docs

link:
	BUILD_DIR=${BUILD_DIR}; \
	BUILD_DIR_LINK=${BUILD_DIR_LINK}; \
	CURRENT_DIR=${CURRENT_DIR}; \
  echo "$${CURRENT_DIR}"; \
  echo "$${BUILD_DIR_LINK}"; \
	if [ "$${BUILD_DIR}" != "$${CURRENT_DIR}" ]; then \
	    echo "Fixing symlinks for build"; \
	    ln -s $${CURRENT_DIR} $${BUILD_DIR}; \
	fi

linux: 
	cd ${BUILD_DIR}; \
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH} . ; \
	cd - >/dev/null

darwin:
	cd ${BUILD_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH} . ; \
	cd - >/dev/null

windows:
	cd ${BUILD_DIR}; \
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe . ; \
	cd - >/dev/null

test:
	if ! hash go2xunit 2>/dev/null; then go install github.com/tebeka/go2xunit; fi
	cd ${BUILD_DIR}; \
	godep go test -v ./... 2>&1 | go2xunit -output ${TEST_REPORT} ; \
	cd - >/dev/null

vet:
	-cd ${BUILD_DIR}; \
	godep go vet ./... > ${VET_REPORT} 2>&1 ; \
	cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${VET_REPORT}
	-rm -f ${BINARY}-*

goget:
	cd ${BUILD_DIR} ; \
	GOBIN= GOOS=linux GOARCH=${GOARCH} go get ; \
	GOBIN= GOOS=darwin GOARCH=${GOARCH} go get ; \
	GOBIN= GOOS=windows GOARCH=${GOARCH} go get ;

docs: linux
	cd ${BUILD_DIR} ; \
	./eva-linux-${GOARCH} docs export markdown ;

.PHONY: link linux darwin windows test vet fmt clean goget docs
