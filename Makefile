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

.PHONY: cfgtest
cfgtest:
	mkdir -p ./testdata/test/config
	mkdir -p ./testdata/test/cache
	WENDOVER_CONFIG=./testdata/test/config WENDOVER_CACHE=./testdata/test/cache go run build/wendover/install.go $(CURDIR)
	WENDOVER_CONFIG=./testdata/test/config WENDOVER_CACHE=./testdata/test/cache go run build/wendsrv/install.go $(CURDIR)

.PHONY: test_db_up
test_db_up:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/test/config go run migrate.go up

.PHONY: test_db_down
test_db_down:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/test/config go run migrate.go down

.PHONY: test_db_reset
test_db_reset:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/test/config go run migrate.go reset

.PHONY: test
test: bin/wendover bin/wendsrv
	$(MAKE) test_db_reset
	go test -v -count=1 ./...
	$(MAKE) test_db_down

.PHONY: install
install: bin/wendover
	go run build/wendover/install.go $(CURDIR)
	cp bin/wendover ${HOME}/.local/bin/

.PHONY: dev_install
dev_install: bin/wendover bin/wendsrv
	mkdir -p ./testdata/dev/config
	mkdir -p ./testdata/dev/cache
	WENDOVER_CONFIG=./testdata/dev/config WENDOVER_CACHE=./testdata/dev/cache go run build/wendover/install.go $(CURDIR)
	WENDOVER_CONFIG=./testdata/dev/config WENDOVER_CACHE=./testdata/dev/cache go run build/wendsrv/install.go $(CURDIR)

.PHONY: dev_clean_install
dev_clean_install: bin/wendover bin/wendsrv
	-rm -r ./testdata/dev/
	$(MAKE) dev_install

.PHONY: dev_srv_run
dev_srv_run: bin/wendsrv
	WENDOVER_CONFIG=./testdata/dev/config bin/wendsrv run

.PHONY: dev_db_up
dev_db_up:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/dev/config go run migrate.go up

.PHONY: dev_db_down
dev_db_down:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/dev/config go run migrate.go down

.PHONY: dev_db_reset
dev_db_reset:
	cd tools/migrate/ && WENDOVER_CONFIG=../../testdata/dev/config go run migrate.go reset

.PHONY: clean
clean:
	rm -rf bin/
