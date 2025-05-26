---
title: "Steampipe Table: pdns_zone - Query pdns Users using SQL"
description: "Allows users to query pdns Users, providing insights into users details such as its full_name, absolute_url and more."
---

# Table: pdns_user – Query pdns zones Using SQL

pdns users are entities configured within a pdns instance, typically representing developers, admins, or service accounts interacting with the pdns system. This table provides insights into the dns records configured in pdns, including their zones and URLs.

## Table Usage Guide
The pdns_zone table allows pdns administrators, DevOps engineers, and auditors to query details about zones in a pdns instance. 

Each row in this table represents a single pdns zone account retrieved from the pdns database or API.

## Examples

### List all zones with URLs
This query retrieves the name and url of all zone configured in your pdns environment. It's useful for getting a complete overview of every zone, including their URLs.

```sql+postgres
select
  name,
  url
from
  pdns_zone;
```

```sql+sqlite
select
  name,
  url
from
  pdns_zone;
```

### List specfic users with URLs
This query filters the list of pdns zone to show only those whose name contains the term "example" (case-insensitive).

```sql+postgres
select
  name,
  url
from
  pdns_zone
where
  name like '%example%';
```

```sql+sqlite
select
  name,
  url
from
  pdns_zone
where
  lower(name) like '%example%';
```