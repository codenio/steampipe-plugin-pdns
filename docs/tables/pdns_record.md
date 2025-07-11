---
title: "Steampipe Table: pdns_record - Query PowerDNS Records using SQL"
description: "Allows users to query PowerDNS Records, providing insights into DNS record configurations and settings."
---

# Table: pdns_record - Query PowerDNS Records Using SQL

PowerDNS Records are individual DNS entries within zones that define how domain names map to IP addresses and other resources. Each record has a specific type (like A, AAAA, MX, etc.) and contains information about how the DNS should respond to queries for that record.

## Table Usage Guide

The `pdns_record` table provides insights into DNS records within PowerDNS zones. As a DNS administrator, you can use this table to query, analyze, and monitor DNS records across your zones. This table is particularly useful when you need to audit your DNS configurations, verify record settings, or troubleshoot DNS-related issues.

## Examples

### List all DNS records
```sql
select
  name,
  type,
  content,
  ttl
from
  pdns_record;
```

### Find all A records with their TTL values
```sql
select
  name,
  content,
  ttl
from
  pdns_record
where
  type = 'A';
```

### List all MX records sorted by priority
```sql
select
  name,
  content,
  priority
from
  pdns_record
where
  type = 'MX'
order by
  priority;
```

### Find records with low TTL values (less than 300 seconds)
```sql
select
  name,
  type,
  ttl
from
  pdns_record
where
  ttl < 300
order by
  ttl;
```

### List disabled records
```sql
select
  name,
  type,
  content
from
  pdns_record
where
  disabled = true;
```

### Count records by type for each zone
```sql
select
  zone_name,
  type,
  count(*) as record_count
from
  pdns_record
group by
  zone_name,
  type
order by
  zone_name,
  record_count desc;
```

### Find records with PTR records enabled
```sql
select
  name,
  type,
  content
from
  pdns_record
where
  set_ptr = true;
``` 