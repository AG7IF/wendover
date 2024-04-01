SRC = $(shell find . -name "*.go")

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=${BUILD_TIME} -X main.GitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.GitBranch=${GIT_BRANCH}"

bin/wendover: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/wendover cmd/wendover/main.go

bin/wendsrv: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/wendsrv cmd/wendsrv/main.go

bin/wendsrv-run-lambda: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -tags lambda.norpc -o bin/wendsrv-run-lambda cmd/wendsrv-run-lambda/main.go

bin/wendsrv-migrate-lambda: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -tags lambda.norpc -o bin/wendsrv-migrate-lambda cmd/wendsrv-migrate-lambda/main.go

.PHONY: install
install: bin/wendover
	go run build/wendover/install.go $(CURDIR)
	cp bin/wendover ${HOME}/.local/bin/

.PHONY: install_server
install_server: bin/wendsrv
	mv bin/wendsrv ${GOPATH}/bin/

.PHONY: build_wendsrv_docker
server:
	docker build -t wendsrv:latest -f deployments/docker/wendsrv/Dockerfile .

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: clean
clean:
	rm -rf bin/
