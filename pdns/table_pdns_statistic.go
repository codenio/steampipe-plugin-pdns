package pdns

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePdnsStatistic() *plugin.Table {
	return &plugin.Table{
		Name:        "pdns_statistic",
		Description: "Statistics from the PowerDNS server.",
		List: &plugin.ListConfig{
			Hydrate: listPdnsStatistics,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: "Name of the statistic."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "Type of the statistic (counter, gauge)."},
			{Name: "value", Type: proto.ColumnType_STRING, Transform: transform.FromField("Value"), Description: "Value of the statistic."},
		},
	}
}

type PDNSStatistic struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func listPdnsStatistics(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := Connect(ctx, d)
	if err != nil {
		return nil, err
	}

	stats, err := client.Statistics.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, stat := range stats {
		if stat.Name == nil || stat.Type == nil {
			continue // Skip invalid statistics
		}

		// Convert the value to string, handling interface{} type
		value := ""
		if stat.Value != nil {
			switch v := stat.Value.(type) {
			case string:
				value = v
			case float64:
				value = fmt.Sprintf("%v", v)
			case int64:
				value = fmt.Sprintf("%d", v)
			case bool:
				value = fmt.Sprintf("%v", v)
			default:
				value = fmt.Sprintf("%v", v)
			}
		}

		d.StreamListItem(ctx, &PDNSStatistic{
			Name:  *stat.Name,
			Type:  *stat.Type,
			Value: value,
		})
	}

	return nil, nil
}
