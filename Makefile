NAME            := software-timeline
OSARCH          := "darwin/amd64 linux/amd64"
GITHUB_USER     := kentakozuka

ifndef GOBIN
GOBIN := $(shell echo "$${GOPATH%%:*}/bin")
endif

LINT := $(GOBIN)/golint
GOX := $(GOBIN)/gox
ARCHIVER := $(GOBIN)/archiver
GHR := $(GOBIN)/ghr

.DEFAULT_GOAL := build

.PHONY: sort-timeline
sort-timeline:
	go run ./cmd/parser/parser.go list ./data/timeline_out.yaml --out=./data/timeline_out.yaml

.PHONY: start-gateway
start-gateway:
	go run ./cmd/gateway/gateway.go start

.PHONY: start-web-client
start-web-client:
	cd react-sample && npm start


# .PHONY: deps
# deps:
#   go get -d -v .

# .PHONY: build
# build: deps
#   go build -ldflags $(LDFLAGS) -o bin/$(NAME)

# .PHONY: install
# install: deps
#   go install -ldflags $(LDFLAGS)

# .PHONY: cross-build
# cross-build: deps $(GOX)
#   rm -rf ./out && \
#   gox -ldflags $(LDFLAGS) -osarch $(OSARCH) -output "./out/${NAME}_${VERSION}_{{.OS}}_{{.Arch}}/{{.Dir}}"

# .PHONY: package
# package: cross-build $(ARCHIVER)
#   rm -rf ./pkg && mkdir ./pkg && \
#   pushd out && \
#   find * -type d -exec archiver make ../pkg/{}.tar.gz {}/$(NAME) \; && \
#   popd

# .PHONY: release
# release: $(GHR)
#   ghr -u $(GITHUB_USER) $(VERSION) pkg/

# .PHONY: lint
# lint: $(LINT)
#   go lint ./...

# .PHONY: vet
# vet:
#   go vet ./...

# .PHONY: test
# test:
#   go test ./...

# .PHONY: check
# check: lint vet test build
