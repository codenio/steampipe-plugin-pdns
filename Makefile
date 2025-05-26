# TODO: Update plugin name from template to actual plugin name
STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/template@latest/steampipe-plugin-aws.plugin -tags "${BUILD_TAGS}" *.go

dev:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/template@latest/steampipe-plugin-aws.plugin -tags "dev ${BUILD_TAGS}" *.go
