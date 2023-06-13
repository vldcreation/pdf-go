clean:
	@echo "cleaning ..."
	@rm -f privy-pdf-go
	@rm -rf vendor
	@rm -f go.sum

install:
	@echo "Installing dependencies...."
	@rm -rf vendor
	@go mod tidy && go mod download && go mod vendor

.PHONY: clean, install