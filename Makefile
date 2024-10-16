CONTAINER_BUILDER=podman
CONTAINER_NAME=phalanx
CONTAINER_VERSION=1.0.0-alpha

APP_BIN=phalanx.exe
SYSTEM_VERSION="1.0.0-alpha"
BUILD_FLAGS:=-trimpath -ldflags="-s -w -X main.VERSION=$(SYSTEM_VERSION) -X main.MODE=Release"

BACKEND=phalanx
BUILD_FILES=./$(phalanx)/cmd

.PHONY: container-build
container-build:
	@$(CONTAINER_BUILDER) build --squash --tag $(CONTAINER_NAME):${CONTAINER_VERSION} .

.PHONY: go-run
go-run:
	@go run $(BUILD_FILES) --env .env

.PHONY: go-fmt
go-fmt:
	@go fmt ./$(phalanx)/...

.PHONY: go-test
go-test:
	@go test ./$(phalanx)/...

.PHONY: go-benchmark
go-benchmark:
	@go test -bench -benchmem  ./$(phalanx)/...

.PHONY: go-credits
go-credits:
	@gocredits --skip-missing . > CREDITS

.PHONY: go-licenses
go-licenses:
	@go-licenses report github.com/google/go-licenses > licenses.csv

.PHONY: go-build
go-build:
	@go build -o bin/$(APP_BIN) $(BUILD_FLAGS) $(BUILD_FILES)
