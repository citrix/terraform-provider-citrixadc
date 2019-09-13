resource "citrixadc_sslcertkey" "generic-cert" {
  certkey            = var.ssl_config["certname"]
  cert               = var.ssl_config["certfile"]
  key                = var.ssl_config["keyfile"]
  expirymonitor      = "ENABLED"
  notificationperiod = var.ssl_config["notificationperiod"]
}

resource "citrixadc_lbvserver" "generic_lb" {
  name            = var.lb_config["lbname"]
  ipv46           = var.lb_config["vip"]
  port            = var.lb_config["port"]
  lbmethod        = "ROUNDROBIN"
  persistencetype = "COOKIEINSERT"
  servicetype     = var.lb_config["servicetype"]
  sslcertkey      = citrixadc_sslcertkey.generic-cert.certkey
  sslprofile      = "ns_default_ssl_profile_secure_frontend"
}

resource "citrixadc_service" "backend" {
  lbvserver   = citrixadc_lbvserver.generic_lb.name
  lbmonitor   = citrixadc_lbmonitor.generic_monitor.monitorname
  count       = length(var.backend_services)
  ip          = element(var.backend_services, count.index)
  servicetype = var.backend_service_config["servicetype"]
  port        = var.backend_service_config["port"]
  clttimeout  = var.backend_service_config["client_timeout"]
}

resource "citrixadc_lbmonitor" "generic_monitor" {
  monitorname = var.http_monitor_config["name"]
  type        = "HTTP"
  interval    = var.http_monitor_config["interval_ms"]
  resptimeout = var.http_monitor_config["response_timeout_ms"]
  units3      = "MSEC"
  units4      = "MSEC"
}
