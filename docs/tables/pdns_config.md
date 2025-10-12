---
title: "Steampipe Table: pdns_config - Query PowerDNS Configuration using SQL"
description: "Allows users to query PowerDNS server configuration settings, providing insights into server configuration parameters and values."
---

# Table: pdns_config - Query PowerDNS Configuration Using SQL

PowerDNS Configuration provides access to the server's configuration settings and parameters that control various aspects of the DNS server operation.

## Table Usage Guide

The `pdns_config` table provides insights into PowerDNS server configuration settings. Use this table to review server configuration, audit settings, and verify server configurations.

## Examples

### List all configuration settings
```sql
select
  name,
  value,
  type
from
  pdns_config
order by
  name;
```

### Find specific configuration settings
```sql
select
  name,
  value,
  type
from
  pdns_config
where
  name like '%cache%'
order by
  name;
```

### List configuration settings by type
```sql
select
  type,
  count(*) as setting_count
from
  pdns_config
group by
  type
order by
  setting_count desc;
```