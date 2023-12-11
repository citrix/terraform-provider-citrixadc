terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
terraform {
  required_version = ">= 0.12"
}

//configure backend load balancing servers
resource "citrixadc_server" "tf_server" {
  for_each  = var.servers
  name      = each.key
  ipaddress = each.value["ipaddress"]
  state     = each.value["state"]
}
//configure load balancing services and parameters
resource "citrixadc_service" "tf_service" {
  for_each  = var.servers
  name         = "${each.key}-${var.service_config["service_port"]}"
  servername   = each.key
  port         = var.service_config["service_port"]
  servicetype  = var.service_config["servicetype"]
  maxclient    = var.service_config["maxclient"]
  maxreq       = var.service_config["maxreq"]
  cip          = var.service_config["cip"]
  cipheader    = var.service_config["cipheader"]
  usip         = var.service_config["usip"]
  useproxyport = var.service_config["useproxyport"]
  sp           = var.service_config["sp"]
  clttimeout   = var.service_config["clttimeout"]
  svrtimeout   = var.service_config["svrtimeout"]
  cka          = var.service_config["cka"]
  tcpb         = var.service_config["tcpb"]
  cmp          = var.service_config["cmp"]
  depends_on   = [citrixadc_server.tf_server]
}