/*
Package gcp implements a steampipe plugin for gcp.

This plugin provides data that Steampipe uses to present foreign
tables that represent GCP resources.
*/
package gcp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/rate_limiter"
)

const pluginName = "steampipe-plugin-gcp"

// Plugin creates this (gcp) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isIgnorableError([]string{"404", "400"}),
		},
		// Default ignore config for the plugin
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrorPluginDefault(),
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "project",
				Hydrate: getProject,
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		RateLimiters: []*rate_limiter.Definition{
			// API Requests per 100 seconds: 5,000
			// https://cloud.google.com/memorystore/docs/redis/quotas#per-second_api_requests_quota
			{
				Name:       "gcp_redis_list_instances",
				FillRate:   50,
				BucketSize: 5000,
				Scope:      []string{"connection", "service", "action"},
				Where:      "service = 'redis' and action = 'ListInstances'",
			},
			{
				Name:       "gcp_redis_get_instance",
				FillRate:   50,
				BucketSize: 5000,
				Scope:      []string{"connection", "service", "action"},
				Where:      "service = 'redis' and action = 'GetInstance'",
			},
			// Redis Cluster requests per project per minute: 60
			// https://cloud.google.com/memorystore/docs/cluster/quotas#per-minute_api_requests_quota
			{
				Name:       "gcp_rediscluster_list_clusters",
				FillRate:   1,
				BucketSize: 60,
				Scope:      []string{"connection", "service", "action"},
				Where:      "service = 'rediscluster' and action = 'ListClusters'",
			},
			{
				Name:       "gcp_rediscluster_get_cluster",
				FillRate:   1,
				BucketSize: 60,
				Scope:      []string{"connection", "service", "action"},
				Where:      "service = 'rediscluster' and action = 'GetCluster'",
			},
		},
		TableMap: map[string]*plugin.Table{
			"template_sample": tableGcpAlloyDBCluster(ctx),
		},
	}

	return p
}
