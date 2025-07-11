---
organization: Turbot
category: ["networking"]
icon_url: "/images/plugins/codenio/pdns.svg"
brand_color: "#F48220"
display_name: "PowerDNS"
name: "pdns"
description: "Steampipe plugin for querying zones, records, servers and more from PowerDNS."
og_description: "Query PowerDNS with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/codenio/pdns-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# PowerDNS + Steampipe

[PowerDNS](https://www.powerdns.com/) is a powerful, open source DNS server software that supports various backends and provides extensive DNS functionality including DNSSEC.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

For example:

```sql
select
  name,
  kind,
  dnssec,
  serial
from
  pdns_zone;
```

```
+---------------+--------+--------+--------+
| name          | kind   | dnssec | serial |
+---------------+--------+--------+--------+
| example.com.  | Native | false  | 123456 |
| test.org.     | Master | true   | 789012 |
| domain.net.   | Slave  | false  | 345678 |
+---------------+--------+--------+--------+
```

## Documentation

- **[Table definitions & examples →](/plugins/codenio/pdns/tables)**

## Quick start

### Install

Download and install the latest PowerDNS plugin:

```bash
steampipe plugin install pdns
```

### Configuration

Connection configuration for PowerDNS Plugin can be accomplished in two ways:

1. In a steampipe config file (`~/.steampipe/config/pdns.spc`):
```hcl
connection "pdns" {
  plugin = "pdns"

  # PowerDNS API server URL (Required)
  # Can also be set with PDNS_URL environment variable
  server_url = "http://localhost:8081"

  # PowerDNS API virtual host (Required)
  # Can also be set with PDNS_VHOST environment variable
  vhost = "localhost"

  # PowerDNS API key (Required)
  # Can also be set with PDNS_APIKEY environment variable
  api_key = "your-api-key-here"
}
```

2. Using environment variables:
```sh
export PDNS_URL=http://localhost:8081
export PDNS_VHOST=localhost
export PDNS_APIKEY=your-api-key-here
```


