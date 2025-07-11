package pdns

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePdnsConfig() *plugin.Table {
	return &plugin.Table{
		Name:        "pdns_config",
		Description: "Configuration settings of the PowerDNS server.",
		List: &plugin.ListConfig{
			Hydrate: listPdnsConfig,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: "Name of the configuration setting."},
			{Name: "value", Type: proto.ColumnType_STRING, Transform: transform.FromField("Value"), Description: "Value of the configuration setting."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "Type of the configuration value."},
		},
	}
}

type PDNSConfig struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

func listPdnsConfig(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := Connect(ctx, d)
	if err != nil {
		return nil, err
	}

	// Get the configuration settings
	config, err := client.Config.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, setting := range config {
		if setting.Name == nil || setting.Value == nil || setting.Type == nil {
			continue // Skip invalid settings
		}

		d.StreamListItem(ctx, &PDNSConfig{
			Name:  *setting.Name,
			Value: *setting.Value,
			Type:  *setting.Type,
		})
	}

	return nil, nil
}
