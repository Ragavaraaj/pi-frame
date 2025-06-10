APP_NAME=photo-frame
GO_FILES=$(wildcard *.go)
OUTPUT_DIR=build
OUTPUT_FILE=$(OUTPUT_DIR)/$(APP_NAME)

.DEFAULT_GOAL := run

.PHONY: fmt vet build brun

fmt:
		@echo "Formatting code..."
		go fmt ./...

vet:
	    @echo "Running go vet..."
	    go vet ./...

build:
		@echo "Creating build directory..."
		mkdir -p $(OUTPUT_DIR)
	    @echo "Building the application..."
	    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags nogpu -o $(OUTPUT_FILE)

brun: build
	    @echo "Running the application in background..."
	    $(OUTPUT_FILE)

run:
	    @echo "Running the application..."
	    SLIDESHOW_INTERVAL=60 SLIDESHOW_DIR=/home/display/images go run main.go
