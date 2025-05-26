package pdns

import (
	"context"
	"errors"
	"os"

	"github.com/joeig/go-powerdns/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*powerdns.Client, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*powerdns.Client), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	pdnsConfig := GetConfig(d.Connection)

	var server_url, vhost, api_key string

	if pdnsConfig.ServerURL != nil {
		server_url = *pdnsConfig.ServerURL
	} else {
		server_url = os.Getenv("PDNS_URL")
	}

	if pdnsConfig.VHost != nil {
		vhost = *pdnsConfig.VHost
	} else {
		vhost = os.Getenv("PDNS_VHOST")
	}

	if pdnsConfig.ApiKey != nil {
		api_key = *pdnsConfig.ApiKey
	} else {
		api_key = os.Getenv("PDNS_APIKEY")
	}

	// Error if the minimum config is not set
	if server_url == "" || vhost == "" || api_key == "" {
		return nil, errors.New("'server_url', 'username' and 'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
	}

	return powerdns.New(server_url, vhost, powerdns.WithAPIKey(api_key)), nil
}
