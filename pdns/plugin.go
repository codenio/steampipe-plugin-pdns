package pdns

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-pdns"

// Plugin creates this (pdns) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"pdns_zone": tablePdnsZone(),
			// "pdns_folder":            tablepdnsFolder(),
			// "pdns_freestyle_project": tablepdnsFreestyleProject(),
			// "pdns_job":               tablepdnsJob(),
			// "pdns_node":              tablepdnsNode(),
			// "pdns_pipeline":          tablepdnsPipeline(),
			// "pdns_plugin":            tablepdnsPlugin(),
			// "pdns_user":              tablepdnsUser(),
		},
	}

	return p
}
