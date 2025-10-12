# TODO: Update plugin name from template to actual plugin name
STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/codnio/pdns@latest/steampipe-plugin-pdns.plugin -tags "${BUILD_TAGS}" *.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/steampipe-plugin-pdns-linux-amd64.plugin -tags "${BUILD_TAGS}" *.go

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/steampipe-plugin-pdns-darwin-amd64.plugin -tags "${BUILD_TAGS}" *.go

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o build/steampipe-plugin-pdns-darwin-arm64.plugin -tags "${BUILD_TAGS}" *.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o build/steampipe-plugin-pdns-windows-amd64.plugin -tags "${BUILD_TAGS}" *.go