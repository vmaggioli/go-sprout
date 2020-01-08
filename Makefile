GO_BUILD=go build
GO_CLEAN=go clean
GO_VERIFY=go mod verify

OUTPUT_DIR=build
OUTPUT_NAME=sprout
SRC_PATH=src/sprout/main.go

WIN_SPROUT=$(OUTPUT_NAME).exe
LINUX_SPROUT=$(OUTPUT_NAME)
MACOS_SPROUT=$(OUTPUT_NAME)

ZIP_WIN=$(OUTPUT_DIR)/windows/sprout_win_x64.tar.gz
ZIP_LINUX=$(OUTPUT_DIR)/linux/$(LINUX_SPROUT)_linux_x64.tar.gz
ZIP_MACOS=$(OUTPUT_DIR)/macos/$(MACOS_SPROUT)_macos_x64.tar.gz

.DEFAULT: all
all:
	$(MAKE) build_base
	$(MAKE) build_win
	$(MAKE) build_mac
	$(MAKE) build_linux

build_base:
	$(GO_BUILD) -o $(OUTPUT_DIR)/$(OUTPUT_NAME) $(SRC_PATH)

build_win:
	$(GO_BUILD) -o $(OUTPUT_DIR)/windows/$(WIN_SPROUT) $(SRC_PATH)
	tar -zcvf $(ZIP_WIN) -C $(OUTPUT_DIR)/windows $(WIN_SPROUT)

build_mac:
	$(GO_BUILD) -o $(OUTPUT_DIR)/macos/$(MACOS_SPROUT) $(SRC_PATH)
	tar -zcvf $(ZIP_MACOS) -C $(OUTPUT_DIR)/macos $(MACOS_SPROUT)

build_linux:
	$(GO_BUILD) -o $(OUTPUT_DIR)/linux/$(LINUX_SPROUT) $(SRC_PATH)
	tar -zcvf $(ZIP_LINUX) -C $(OUTPUT_DIR)/linux $(LINUX_SPROUT)

verify_build:
	$(GOVERIFY)
	
verify_win: verify_build build_win

verify_linux: verify_build build_linux

verify_mac: verify_build build_mac

clean:
	$(GO_CLEAN)

	# Executables
	rm -f $(OUTPUT_DIR)/windows/$(WIN_SPROUT)
	rm -f $(OUTPUT_DIR)/macos/$(MACOS_SPROUT)
	rm -f $(OUTPUT_DIR)/linux/$(MACOS_SPROUT)

	# Zips
	rm -f $(ZIP_WIN)
	rm -f $(ZIP_MACOS)
	rm -f $(ZIP_LINUX)