APP_NAME=fbimage_display
GO_FILES=$(wildcard *.go)
OUTPUT_DIR=build
OUTPUT_FILE=$(OUTPUT_DIR)/$(APP_NAME)

.PHONY: all build clean test 

all: build

build: $(GO_FILES)
	@mkdir -p $(OUTPUT_DIR)
	@echo "Building executable..."
	@GOOS=linux GOARCH=arm GOARM=6 go build -o $(OUTPUT_FILE)

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(OUTPUT_DIR)


ci: test build
	@echo "CI pipeline completed successfully."