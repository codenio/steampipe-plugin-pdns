package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-template/template"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: template.Plugin})
}
