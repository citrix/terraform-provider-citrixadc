lb_config = {
  lbname      = "ProductionLB"
  vip         = "10.22.24.22"
  servicetype = "SSL"
  port        = 443
}

backend_service_config_cart = {
  name           = "cart"
  url            = "/cart/*"
  port           = 8080
  servicetype    = "HTTP"
  client_timeout = 30
}

backend_service_config_catalog = {
  name           = "catalog"
  url            = "/catalog/*"
  port           = 8000
  servicetype    = "HTTP"
  client_timeout = 90
}

backend_services_cart = [
  "172.25.33.53:8080",
  "172.25.44.53:8080",
  "172.25.44.54:8080",
]

backend_services_catalog = [
  "172.23.33.33:8000",
  "172.23.44.33:8000",
]

ssl_config = {
  certname           = "top_secret"
  certfile           = "/var/certs/server.crt"
  keyfile            = "/var/certs/server.key"
  notificationperiod = 90
}

http_monitor_config_cart = {
  name                = "down_monitor_cart"
  interval_ms         = 150
  response_timeout_ms = 50
}

http_monitor_config_catalog = {
  name                = "down_monitor_catalog"
  interval_ms         = 250
  response_timeout_ms = 60
}
