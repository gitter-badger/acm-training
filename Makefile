GO := go
GOGET := $(GO) get
GOBUILD := $(GO) build
GOBUILD_FLAGS := -v
GOTEST := $(GO) test
GOTEST_FLAGS := -v
TARGET := acm-training
GOLINT := golint

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
	@which golint &> /dev/null || $(GOGET) -u golang.org/x/lint/golint

.PHONY: lint
lint: $(TEST_SOURCES) .golint_binary_check
	@echo "Perform linting..."
	@$(GOLINT) ./...

.PHONY: test
test: $(TEST_SOURCES) lint
	@echo "Perform testing..."
	@$(GOTEST) $(GOTEST_FLAGS) ./...

.PHONY: .go_acc_binary_check
.go_acc_binary_check:
	@echo "Checking go-acc toolchain..."
	@which go-acc &> /dev/null || $(GOGET) -u github.com/ory/go-acc

.PHONY: coverage
coverage: $(TEST_SOURCES) lint .go_acc_binary_check
	@echo "Generating coverage report..."
	@go-acc github.com/uestc-acm/acm-training/...

.PHONY: clean
clean:
	@echo "Perform cleanup..."
	@rm -f $(TARGET) $(COVERAGE_REPORT)
