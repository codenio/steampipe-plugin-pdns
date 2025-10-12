connection "pdns" {
  plugin = "ghcr.io/codenio/pdns"

  # PowerDNS API server URL (Required)
  # Can also be set with PDNS_URL environment variable
  # server_url = "http://localhost:8081"

  # PowerDNS API virtual host (Required)
  # Can also be set with PDNS_VHOST environment variable
  # vhost = "localhost"

  # PowerDNS API key (Required)
  # Can also be set with PDNS_APIKEY environment variable
  # api_key = "ODl4Tkx2c3erZnY4W***"

  # For more information about PowerDNS API configuration:
  # https://doc.powerdns.com/authoritative/http-api/index.html
}
