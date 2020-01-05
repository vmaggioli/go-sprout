SRC_PATH=src/sprout/main.go

GO_BUILD=go build
GO_CLEAN=go clean

OUTPUT_DIR=build
OUTPUT_NAME=sprout

.DEFAULT: all
all:
	$(MAKE) build_base

build_base:
	$(GO_BUILD) -o $(OUTPUT_DIR)/$(OUTPUT_NAME) $(SRC_PATH)

clean:
	$(GO_CLEAN)
	rm -f $(OUTPUT_DIR)/$(OUTPUT_NAME)