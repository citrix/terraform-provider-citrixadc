lb_config = {
  lbname      = "AuctionLB"
  vip         = "10.22.24.22"
  servicetype = "SSL"
  port        = 443
}

backend_service_config = {
  servicetype    = "HTTP"
  port           = 8080
  client_timeout = 30
}

backend_services = [
  "172.23.33.33",
  "172.23.44.33",
  "172.23.44.34",
]

ssl_config = {
  certname           = "top_secret"
  certfile           = "/var/certs/server.crt"
  keyfile            = "/var/certs/server.key"
  notificationperiod = 83
}

http_monitor_config = {
  name                = "down_monitor"
  interval_ms         = 150
  response_timeout_ms = 50
}
