package pdns

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type pdnsConfig struct {
	ServerURL *string `hcl:"server_url"`
	VHost     *string `hcl:"vhost"`
	ApiKey    *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &pdnsConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) pdnsConfig {
	if connection == nil || connection.Config == nil {
		return pdnsConfig{}
	}
	config, _ := connection.Config.(pdnsConfig)
	return config
}
