include .env
export

build_db:
	@echo "Building MongoDB Container..."
	@docker-compose up -d || (echo "Failed to build MongoDB container" && exit 1)
	@echo.
	@echo "MongoDB running on localhost:8081"

run_mangomarkets:
	@echo.
	@echo "Launching MangoMarkets..."
	@echo.
	@go build -o ./bin/mangomarkets.exe ./cmd/mangomarkets || (echo "Failed to build MangoMarkets" && exit 1)
	@.\bin\mangomarkets.exe || echo "Failed to run MangoMarkets"

all: build_db run_mangomarkets