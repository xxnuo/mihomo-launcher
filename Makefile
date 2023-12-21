# build launcher

NAME=mihomo-launcher
BINDIR=bin
BRANCH=$(shell cd ../core && git branch --show-current)
VERSION=$(shell cd ../core && git rev-parse --short HEAD)

BUILDTIME=$(shell date -u)
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '\
		-X "github.com/xnuo/mihomo-launcher/cmd/mihomo-launcher/constants.Version=$(VERSION)" \
		-X "github.com/metacubex/mihomo/constant.BuildTime=$(BUILDTIME)" \
		-w -s -H windowsgui -buildid='

USE_UPX=1

windows-amd64:
	GOARCH=amd64 GOOS=windows GOAMD64=v3 $(GOBUILD) -o ./$(BINDIR)/$(NAME)-$@.exe ./cmd/mihomo-launcher
ifeq ($(USE_UPX),1)
	upx ./$(BINDIR)/$(NAME)-$@.exe
endif




# 更新所需文件 FOR WINDOWS

OUT_BIN_DIR=cfs/root/bin

CORE_DIR=../core
CORE_NAME=mihomo
CORE_BUILD_DIR=bin
CORE_PLATFORM=windows-amd64

update-core: $(CORE_DIR)/Makefile
	@echo "Building core..."
	@rm -rf ./$(CORE_DIR)/$(CORE_BUILD_DIR)
	@cd ./$(CORE_DIR) && git pull --force origin Alpha && make $(CORE_PLATFORM)
	@mkdir -p ./$(OUT_BIN_DIR)
	@rm -rf ./$(OUT_BIN_DIR)/$(CORE_NAME).exe
	@upx ./$(CORE_DIR)/$(CORE_BUILD_DIR)/$(CORE_NAME)-$(CORE_PLATFORM).exe -o ./$(OUT_BIN_DIR)/$(CORE_NAME).exe

UI_DIR=../ui
UI_BUILD_DIR=dist
OUT_UI_DIR=cfs/root/ui

update-ui: $(UI_DIR)/package.json
	@echo "Building ui..."
	@rm -rf ./$(UI_DIR)/$(UI_BUILD_DIR)
	@cd ./$(UI_DIR) && git pull --force && pnpm install && pnpm run build
	@mkdir -p ./$(OUT_ROOT)/$(OUT_UI_DIR)
	@rm -rf ./$(OUT_UI_DIR)/*
	@cp -r ./$(UI_DIR)/$(UI_BUILD_DIR)/* ./$(OUT_UI_DIR)
