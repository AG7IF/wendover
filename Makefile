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

.PHONY: server
server: bin/wendsrv

.PHONY: install
install: bin/wendover
	go run build/wendover/install.go $(CURDIR)
	cp bin/wendover ${HOME}/.local/bin/

.PHONY: dev_install
dev_install: bin/wendover bin/wendsrv
	mkdir -p ./testdata/config
	mkdir -p ./testdata/cache
	WENDOVER_CONFIG=./testdata/config WENDOVER_CACHE=./testdata/cache go run build/wendover/install.go $(CURDIR)
	WENDOVER_CONFIG=./testdata/config WENDOVER_CACHE=./testdata/cache go run build/wendsrv/install.go $(CURDIR)

.PHONY: dev_clean_install
dev_clean_install: bin/wendover bin/wendsrv
	-rm -r ./testdata
	$(MAKE) dev_install

.PHONY: dev_srv_run
dev_srv_run: bin/wendsrv
	WENDOVER_CONFIG=./testdata/config bin/wendsrv run

.PHONY: dev_db_up
dev_db_up:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/config go run migrate.go up

.PHONY: dev_db_down
dev_db_down:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/config go run migrate.go down

.PHONY: dev_db_reset
dev_db_reset:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/config go run migrate.go reset

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: clean
clean:
	rm -rf bin/
