---
title: "Steampipe Table: pdns_statistic - Query PowerDNS Statistics using SQL"
description: "Allows users to query PowerDNS server statistics, providing insights into server performance and operational metrics."
---

# Table: pdns_statistic - Query PowerDNS Statistics Using SQL

PowerDNS Statistics provide real-time metrics about the operation of your PowerDNS server. These statistics include query counts, cache performance, latency metrics, and various other operational indicators that help monitor and maintain the health of your DNS infrastructure.

## Table Usage Guide

The `pdns_statistic` table provides insights into PowerDNS server performance metrics. As a system administrator or DNS operator, you can use this table to monitor server health, analyze query patterns, and identify potential performance issues. This table is particularly useful when you need to track server performance, analyze traffic patterns, or troubleshoot performance-related issues.

## Examples

### List all statistics
```sql
select
  name,
  type,
  value
from
  pdns_statistic
order by
  name;
```

### Get query-related statistics
```sql
select
  name,
  value,
  type
from
  pdns_statistic
where
  name like '%query%'
order by
  name;
```

### Find statistics with high values (potential areas of concern)
```sql
select
  name,
  value,
  type
from
  pdns_statistic
where
  cast(value as decimal) > 1000000
order by
  cast(value as decimal) desc;
```

### List all counter-type statistics
```sql
select
  name,
  value
from
  pdns_statistic
where
  type = 'counter'
order by
  name;
```

### Get cache-related statistics
```sql
select
  name,
  value,
  type
from
  pdns_statistic
where
  name like '%cache%'
order by
  name;
```

### Monitor error-related statistics
```sql
select
  name,
  value,
  type
from
  pdns_statistic
where
  name like '%error%'
  or name like '%fail%'
order by
  cast(value as decimal) desc;
```

### Get latency-related statistics
```sql
select
  name,
  value,
  type
from
  pdns_statistic
where
  name like '%latency%'
order by
  name;
``` 