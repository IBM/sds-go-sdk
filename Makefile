# Makefile to build the project
GO=go
LINT=golangci-lint
LINTOPTS=
TEST_TAGS=
COVERAGE=-coverprofile=coverage.txt -covermode=atomic

all: tidy test lint
travis-ci: tidy test-cov lint

test:
	${GO} test ./... ${TEST_TAGS}

test-cov:
	${GO} test ./... ${TEST_TAGS} ${COVERAGE}

test-int-cov:
	${GO} test ./... -tags=integration ${COVERAGE}

test-examples-block:
	go test ./sdsaasv1/... -tags=examples -ginkgo.label-filter="!object" -v

test-examples-object:
	go test ./sdsaasv1/... -tags=examples -ginkgo.label-filter="!block" -v

test-integration-block:
	go test ./sdsaasv1/... -tags=integration -ginkgo.label-filter="!object" -v

test-integration-object:
	go test ./sdsaasv1/... -tags=integration -ginkgo.label-filter="!block" -v

lint:
	${LINT} run --build-tags=integration,examples ${LINTOPTS}

tidy:
	${GO} mod tidy
