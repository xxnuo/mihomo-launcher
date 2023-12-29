CORE_DIR=../core
CORE_NAME=mihomo
CORE_BUILD_DIR=bin
CORE_PLATFORM=windows-amd64

UI_DIR=../ui
UI_BUILD_DIR=dist

OUT_DIR=cfs/root
OUT_BIN_DIR=bin
OUT_UI_DIR=ui

update-core: $(CORE_DIR)/Makefile
	@echo "Building core..."
	@rm -rf ./$(CORE_DIR)/$(CORE_BUILD_DIR)
	@cd ./$(CORE_DIR) && git pull --force origin Alpha && make $(CORE_PLATFORM)
	@mkdir -p ./$(OUT_DIR)/$(OUT_BIN_DIR)
	@rm -rf ./$(OUT_DIR)/$(OUT_BIN_DIR)/$(CORE_NAME).exe
	@upx ./$(CORE_DIR)/$(CORE_BUILD_DIR)/$(CORE_NAME)-$(CORE_PLATFORM).exe -o ./$(OUT_DIR)/$(OUT_BIN_DIR)/$(CORE_NAME).exe

update-ui: $(UI_DIR)/package.json
	@echo "Building ui..."
	@rm -rf ./$(UI_DIR)/$(UI_BUILD_DIR)
	@cd ./$(UI_DIR) && git pull --force && pnpm install && pnpm run build
	@mkdir -p ./$(OUT_DIR)/$(OUT_UI_DIR)
	@rm -rf ./$(OUT_DIR)/$(OUT_UI_DIR)/*
	@cp -r ./$(UI_DIR)/$(UI_BUILD_DIR)/* ./$(OUT_DIR)/$(OUT_UI_DIR)

update-res:
	@echo "Updating resources..."
	@echo "Download very slow, if it fails please try again."
	@mkdir -p ./$(OUT_DIR)/$(OUT_DATA_DIR)
	@curl -o ./$(OUT_DIR)/$(OUT_DATA_DIR)/geoip.dat "https://mirror.ghproxy.com/https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geoip.dat"
	@curl -o ./$(OUT_DIR)/$(OUT_DATA_DIR)/geosite.dat "https://mirror.ghproxy.com/https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geosite.dat"
	@curl -o ./$(OUT_DIR)/$(OUT_DATA_DIR)/country.mmdb "https://mirror.ghproxy.com/https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/country.mmdb"

# update done

launcher:
# go build -ldflags "-H windowsgui -w -s" .\src\cmd\mihomo-launcher
	@echo "Building launcher..."
	@go build -ldflags "-H windowsgui -w -s" -o ./mihomo-launcher.exe ./src/cmd/mihomo-launcher
	@upx ./mihomo-launcher.exe