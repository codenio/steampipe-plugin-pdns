<!-- ![image](https://hub.steampipe.io/images/plugins/turbot/aws-social-graphic.png) -->

# PowerDNS Plugin for Steampipe

Use SQL to query zones, records, statistics and more from PowerDNS.

- **[Get started →](https://hub.steampipe.io/plugins/codenio/pdns)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/codenio/pdns/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/codenio/steampipe-plugin-pdns/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install pdns
```

Configure the PowerDNS connection:

```shell
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/pdns.spc
```

Run a query:

```sql
select 
  name,
  kind,
  dnssec,
  serial 
from 
  pdns_zone;
```

## Available Tables

- `pdns_zone` - DNS zones managed by PowerDNS
- `pdns_record` - DNS records within zones
- `pdns_statistic` - PowerDNS server statistics
- `pdns_config` - PowerDNS server configuration

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/codenio/steampipe-plugin-pdns.git
cd steampipe-plugin-pdns
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/pdns.spc
```

Try it!

```
steampipe query
> .inspect pdns
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [PowerDNS Plugin](https://github.com/codenio/steampipe-plugin-pdns/labels/help%20wanted)
