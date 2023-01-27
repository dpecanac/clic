PROJECT_NAME := "clic"

.PHONY: all
all: help

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the project
	@echo "Building $(PROJECT_NAME)..."
	@go mod tidy
	@go build -o bin/$(PROJECT_NAME) main.go util.go runner.go data.go


clean: ## Clean the project
	@echo "Cleaning $(PROJECT_NAME)..."
	@rm -rf bin

run: clean ## Run the project
	@echo "Running $(PROJECT_NAME)..."
	@./bin/$(PROJECT_NAME)
	@echo "Done!"
