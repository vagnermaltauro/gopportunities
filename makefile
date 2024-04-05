.PHONY: default run build rest docs clean

APP_NAME=gopportunities

default: run

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test -v ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
