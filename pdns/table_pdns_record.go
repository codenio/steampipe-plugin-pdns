package pdns

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePdnsRecord() *plugin.Table {
	return &plugin.Table{
		Name:        "pdns_record",
		Description: "DNS records in PowerDNS zones.",
		List: &plugin.ListConfig{
			ParentHydrate: listPdnsZone,
			Hydrate:       listPdnsRecords,
		},
		Columns: []*plugin.Column{
			{Name: "zone_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ZoneID"), Description: "The ID of the zone this record belongs to."},
			{Name: "zone_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("ZoneName"), Description: "The name of the zone this record belongs to."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: "The name of the record (FQDN)."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "The type of the record (A, AAAA, MX, etc.)."},
			{Name: "ttl", Type: proto.ColumnType_INT, Transform: transform.FromField("TTL"), Description: "Time to live value for the record."},
			{Name: "content", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content"), Description: "The content/value of the record."},
			{Name: "disabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Disabled"), Description: "Whether the record is disabled."},
			{Name: "set_ptr", Type: proto.ColumnType_BOOL, Transform: transform.FromField("SetPTR"), Description: "Whether a PTR record should be created."},
			{Name: "priority", Type: proto.ColumnType_INT, Transform: transform.FromField("Priority"), Description: "Priority of the record (used for MX and SRV records)."},
		},
	}
}

type PDNSRecord struct {
	ZoneID   string `json:"zone_id"`
	ZoneName string `json:"zone_name"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	TTL      int    `json:"ttl"`
	Content  string `json:"content"`
	Disabled bool   `json:"disabled"`
	SetPTR   bool   `json:"set_ptr"`
	Priority int    `json:"priority"`
}

func listPdnsRecords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	zone := h.Item.(map[string]interface{})

	client, err := Connect(ctx, d)
	if err != nil {
		return nil, err
	}

	// Get zone name and ID, handling potential nil values
	var zoneName, zoneID string
	if name, ok := zone["Name"].(*string); ok && name != nil {
		zoneName = *name
	}
	if id, ok := zone["ID"].(*string); ok && id != nil {
		zoneID = *id
	}

	rrsets, err := client.Records.Get(ctx, zoneName, "", nil)
	if err != nil {
		return nil, err
	}

	for _, rrset := range rrsets {
		if rrset.Name == nil || rrset.Type == nil || rrset.TTL == nil {
			continue // Skip records with missing required fields
		}

		for _, record := range rrset.Records {
			if record.Content == nil {
				continue // Skip records with missing content
			}

			// Set default values for optional fields
			disabled := false
			if record.Disabled != nil {
				disabled = *record.Disabled
			}

			setPTR := false
			if record.SetPTR != nil {
				setPTR = *record.SetPTR
			}

			d.StreamListItem(ctx, &PDNSRecord{
				ZoneID:   zoneID,
				ZoneName: zoneName,
				Name:     *rrset.Name,
				Type:     string(*rrset.Type),
				TTL:      int(*rrset.TTL),
				Content:  *record.Content,
				Disabled: disabled,
				SetPTR:   setPTR,
				Priority: 0, // Priority is not directly available in the API response
			})
		}
	}

	return nil, nil
}
