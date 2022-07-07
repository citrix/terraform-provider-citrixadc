resource "citrixadc_sslcertkey" "generic-cert" {
  certkey            = var.ssl_config["certname"]
  cert               = var.ssl_config["certfile"]
  key                = var.ssl_config["keyfile"]
  expirymonitor      = "ENABLED"
  notificationperiod = var.ssl_config["notificationperiod"]
}

resource "citrixadc_csvserver" "generic_cs" {
  name        = var.lb_config["lbname"]
  ipv46       = var.lb_config["vip"]
  port        = var.lb_config["port"]
  servicetype = var.lb_config["servicetype"]
  sslcertkey  = citrixadc_sslcertkey.generic-cert.certkey
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_cspolicy" "cart" {
  policyname      = "cart_cspolicy"
  url             = var.backend_service_config_cart["url"]
  csvserver       = citrixadc_csvserver.generic_cs.name
  targetlbvserver = citrixadc_lbvserver.lb_cart.name
}

resource "citrixadc_cspolicy" "catalog" {
  policyname      = "catalog_cspolicy"
  url             = var.backend_service_config_catalog["url"]
  csvserver       = citrixadc_csvserver.generic_cs.name
  targetlbvserver = citrixadc_lbvserver.lb_catalog.name
}

resource "citrixadc_lbvserver" "lb_cart" {
  name            = var.backend_service_config_cart["name"]
  lbmethod        = "ROUNDROBIN"
  persistencetype = "COOKIEINSERT"
  servicetype     = var.backend_service_config_cart["servicetype"]
}

resource "citrixadc_lbvserver" "lb_catalog" {
  name        = var.backend_service_config_catalog["name"]
  lbmethod    = "LEASTRESPONSETIME"
  servicetype = var.backend_service_config_catalog["servicetype"]
}

resource "citrixadc_servicegroup" "backend_cart" {
  servicegroupname    = var.backend_service_config_cart["name"]
  lbvservers          = [citrixadc_lbvserver.lb_cart.name]
  lbmonitor           = citrixadc_lbmonitor.cart_monitor.monitorname
  servicetype         = var.backend_service_config_cart["servicetype"]
  clttimeout          = var.backend_service_config_cart["client_timeout"]
  servicegroupmembers = var.backend_services_cart
}

resource "citrixadc_servicegroup" "backend_catalog" {
  servicegroupname    = var.backend_service_config_catalog["name"]
  lbvservers          = [citrixadc_lbvserver.lb_catalog.name]
  lbmonitor           = citrixadc_lbmonitor.catalog_monitor.monitorname
  servicegroupmembers = var.backend_services_catalog
  servicetype         = var.backend_service_config_catalog["servicetype"]
  clttimeout          = var.backend_service_config_catalog["client_timeout"]
}

resource "citrixadc_lbmonitor" "cart_monitor" {
  monitorname = var.http_monitor_config_cart["name"]
  type        = "HTTP"
  interval    = var.http_monitor_config_cart["interval_ms"]
  resptimeout = var.http_monitor_config_cart["response_timeout_ms"]
  units3      = "MSEC"
  units4      = "MSEC"
}

resource "citrixadc_lbmonitor" "catalog_monitor" {
  monitorname = var.http_monitor_config_catalog["name"]
  type        = "HTTP"
  interval    = var.http_monitor_config_catalog["interval_ms"]
  resptimeout = var.http_monitor_config_catalog["response_timeout_ms"]
  units3      = "MSEC"
  units4      = "MSEC"
}
