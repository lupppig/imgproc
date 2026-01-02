APP_NAME := imgproc
CMD_DIR := cmd/imgproc
BUILD_DIR := bin

GO := go
GOFLAGS := -trimpath
LDFLAGS := -s -w

.PHONY: all build clean run fmt vet test

all: build

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" \
		-o $(BUILD_DIR)/$(APP_NAME) \
		./$(CMD_DIR)

run:
	$(GO) run ./$(CMD_DIR)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test ./...

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" \
		-o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 \
		./$(CMD_DIR)

# build-darwin:
# 	GOOS=darwin GOARCH=arm64 $(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" \
# 		-o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 \
# 		./$(CMD_DIR)

# build-windows:
# 	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" \
# 		-o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe \
# 		./$(CMD_DIR)
