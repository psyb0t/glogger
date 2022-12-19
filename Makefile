PKG := "github.com/psyb0t/glogger"
PKG_LIST := $(shell go list $(PKG)/...)

dep: ## Get the dependencies + remove unused ones
	@go mod tidy
	@go mod download

test: ## Run tests
	@go test -v $(PKG_LIST)

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
