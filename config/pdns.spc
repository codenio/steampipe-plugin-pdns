connection "pdns" {
  plugin = "pdns"

  # The pdns server URL is required for all requests. Required.
  # It should be fully qualified (e.g. # https://...) and point to the root of the pdns server location.
  # Can also be set via the PDNS_URL environment variable.
  # server_url = "https://ci-cd.internal.my-company.com"

  # The pdns username for authentication is required for requests. Required.
  # Can also be set via the PDNS_VHOST environment variable.
  # vhost = "localhost"

  # Either the pdns password or the API token is required for requests. Required. 
  # Can also be set via the PDNS_API_KEY environment variable.
  # api_key = "aqt*abc8vcf9abc.ABC"

  # Further information: https://www.pdns.io/doc/book/using/using-credentials/   
}
