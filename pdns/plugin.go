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
			"pdns_zone":      tablePdnsZone(),
			"pdns_record":    tablePdnsRecord(),
			"pdns_statistic": tablePdnsStatistic(),
			"pdns_config":    tablePdnsConfig(),
		},
	}

	return p
}
