
resource "netscaler_sslcertkey" "generic-cert" {
  certkey = "${lookup(var.ssl_config, "certname")}"
  cert = "${lookup(var.ssl_config, "certfile")}"
  key = "${lookup(var.ssl_config, "keyfile")}"
  expirymonitor = "ENABLED"
  notificationperiod = "${lookup(var.ssl_config, "notificationperiod")}"
}

resource "netscaler_lbvserver" "generic_lb" {
  name = "${lookup(var.lb_config, "lbname")}"
  ipv46 = "${lookup(var.lb_config, "vip")}"
  port = "${lookup(var.lb_config, "port")}"
  lbmethod = "ROUNDROBIN"
  persistencetype = "COOKIEINSERT"
  servicetype = "${lookup(var.lb_config, "servicetype")}"
  sslcertkey = "${netscaler_sslcertkey.generic-cert.certkey}"
}


resource "netscaler_service" "backend" {
  lbvserver = "${netscaler_lbvserver.generic_lb.name}"
  lbmonitor = "${netscaler_lbmonitor.generic_monitor.monitorname}"
  count = "${length(var.backend_services)}"
  ip = "${element(var.backend_services, count.index)}"
  servicetype = "${lookup(var.backend_service_config, "servicetype")}"
  port = "${lookup(var.backend_service_config, "port")}"
  clttimeout = "${lookup(var.backend_service_config, "client_timeout")}"
}

resource "netscaler_lbmonitor" "generic_monitor" {
  monitorname = "${lookup(var.http_monitor_config, "name")}"
  type = "HTTP"
  interval = "${lookup(var.http_monitor_config, "interval_ms")}"
  resptimeout = "${lookup(var.http_monitor_config, "response_timeout_ms")}"
  units3 = "MSEC"
  units4 = "MSEC"
}
