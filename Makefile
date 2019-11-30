GO := go
GOGET := $(GO) get
GOBUILD := $(GO) build
GOBUILD_FLAGS :=
GOTEST := $(GO) test
GOTEST_FLAGS :=
TARGET := acm-training
GOLINT := golint
GOCLEAN := $(GO) clean

COVERAGE_REPORT := coverage.txt
TEST_SOURCES := $(wildcard **/*.go) main.go
SOURCES := $(filter-out $(wilcard **/*_test.go), $(TEST_SOURCES))

.PHONY: all
all: $(TARGET)

$(TARGET): $(SOURCES)
	@echo "Building the application..."
	@$(GOBUILD) $(GOBUILD_FLAGS) ./...
	@$(GOBUILD) -o $@ $(GOBUILD_FLAGS)

.PHONY: .golint_binary_check
.golint_binary_check:
	@echo "Checking golint toolchain..."
	@which golint > /dev/null || $(GOGET) -u golang.org/x/lint/golint

.PHONY: lint
lint: .golint_binary_check
	@echo "Perform linting..."
	@$(GOLINT) ./...

.PHONY: test
test: $(TEST_SOURCES)
	@echo "Perform testing..."
	@$(GOTEST) $(GOTEST_FLAGS) ./...

.PHONY: coverage
coverage: test lint
	@echo "Generating coverage report..."
	@$(GOTEST) -race -coverprofile=$(COVERAGE_REPORT) -covermode=atomic

.PHONY: clean
clean:
	@echo "Perform cleanup..."
	@$(GOCLEAN)
	@rm -f $(TARGET)
