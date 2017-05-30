
resource "netscaler_lbvserver" "production_lb" {
  name = "productionLB"
  ipv46 = "${lookup(var.vip_config, "vip")}"
  port = "80"
  servicetype = "HTTP"
}

resource "netscaler_servicegroup" "backend" {
  servicegroupname = "productionBackend"
  lbvserver = "${netscaler_lbvserver.production_lb.name}"
  servicetype = "HTTP"
  clttimeout = "${lookup(var.backend_service_config, "clttimeout")}"
  servicegroupmembers = "${formatlist("%s:%s", var.backend_services, var.backend_service_config["backend_port"])}"
}

