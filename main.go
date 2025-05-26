package main

import (
	"github.com/codenio/steampipe-plugin-pdns/pdns"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: pdns.Plugin})
}
