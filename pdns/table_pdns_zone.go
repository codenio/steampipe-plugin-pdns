package pdns

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION
/*
   "account": "",
   "catalog": "",
   "dnssec": false,
   "edited_serial": 1,
   "id": "example.com.",
   "kind": "Native",
   "last_check": 0,
   "masters": [],
   "name": "example.com.",
   "notified_serial": 0,
   "serial": 1,
   "url": "/api/v1/servers/localhost/zones/example.com."
*/

func tablePdnsZone() *plugin.Table {
	return &plugin.Table{
		Name:        "pdns_zone",
		Description: "An extension to pdns functionality provided separately from pdns Core.",
		List: &plugin.ListConfig{
			Hydrate: listPdnsZone,
		},

		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Hydrate: listPdnsZone, Transform: transform.FromField("Name"), Description: "The DNS zone name (FQDN with trailing dot)."},
			{Name: "id", Type: proto.ColumnType_STRING, Hydrate: listPdnsZone, Transform: transform.FromField("ID"), Description: "Internal unique identifier for the zone (typically same as name)."},
			{Name: "kind", Type: proto.ColumnType_STRING, Hydrate: listPdnsZone, Transform: transform.FromField("Kind"), Description: "The type of zone: Native, Master, or Slave."},
			{Name: "account", Type: proto.ColumnType_STRING, Hydrate: listPdnsZone, Transform: transform.FromField("Account"), Description: "Optional account name for multi-tenant setups."},
			{Name: "catalog", Type: proto.ColumnType_STRING, Hydrate: listPdnsZone, Transform: transform.FromField("Catalog"), Description: "Catalog zone name if the zone is part of a DNS catalog."},
			{Name: "dnssec", Type: proto.ColumnType_BOOL, Hydrate: listPdnsZone, Transform: transform.FromField("DNSsec"), Description: "Whether DNSSEC is enabled for the zone."},
			{Name: "serial", Type: proto.ColumnType_INT, Hydrate: listPdnsZone, Transform: transform.FromField("Serial"), Description: "The current serial number of the SOA record."},
			{Name: "edited_serial", Type: proto.ColumnType_INT, Hydrate: listPdnsZone, Transform: transform.FromField("EditedSerial"), Description: "Serial number set manually or most recently edited."},
			{Name: "notified_serial", Type: proto.ColumnType_INT, Hydrate: listPdnsZone, Transform: transform.FromField("NotifiedSerial"), Description: "Last serial number notified to slaves."},
			{Name: "masters", Type: proto.ColumnType_JSON, Hydrate: listPdnsZone, Transform: transform.FromField("Masters"), Description: "List of master server IPs for slave zones."},
			{Name: "url", Type: proto.ColumnType_STRING, Hydrate: listPdnsZone, Transform: transform.FromField("URL"), Description: "API URL to access this zone."},
		},
	}
}

//// LIST FUNCTION

func listPdnsZone(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := Connect(ctx, d)
	if err != nil {
		return nil, err
	}
	zones, err := client.Zones.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		d.StreamListItem(ctx, map[string]interface{}{
			"Name":           zone.Name,
			"ID":             zone.ID,
			"Kind":           zone.Kind,
			"Account":        zone.Account,
			"Catalog":        zone.Catalog,
			"DNSsec":         zone.DNSsec,
			"Serial":         zone.Serial,
			"EditedSerial":   zone.EditedSerial,
			"NotifiedSerial": zone.NotifiedSerial,
			"Masters":        zone.Masters,
			"URL":            zone.URL,
		})
	}
	return nil, nil
}
