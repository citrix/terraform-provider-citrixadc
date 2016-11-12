
resource "netscaler_lbvserver" "generic_lb" {
  name = "${lookup(var.lb_config, "lbname")}"
  ipv46 = "${lookup(var.lb_config, "vip")}"
  port = "${lookup(var.lb_config, "port")}"
  servicetype = "${lookup(var.lb_config, "servicetype")}"
}

resource "netscaler_service" "backend" {
  lbvserver = "${netscaler_lbvserver.generic_lb.name}"
  count = "${length(var.backend_services)}"
  ip = "${element(var.backend_services, count.index)}"
  servicetype = "${lookup(var.backend_service_config, "servicetype")}"
  port = "${lookup(var.backend_service_config, "port")}"
}

