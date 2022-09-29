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
resource "citrixadc_lbvserver" "tf_lbvserverssl" {
  name            = var.sslvip_config["name"]
  ipv46           = var.sslvip_config["vip"]
  port            = var.sslvip_config["port"]
  servicetype     = var.sslvip_config["servicetype"]
  lbmethod        = var.sslvip_config["lbmethod"]
  persistencetype = var.sslvip_config["persistencetype"]
  timeout         = var.sslvip_config["timeout"]
  httpprofilename = var.sslvip_config["httpprofile"]
  redirectfromport = var.sslvip_config["redirectfromport"]
  httpsredirecturl = var.sslvip_config["httpsredirecturl"]
}
//configure SSL vServer settings
resource "citrixadc_sslvserver" "tf_sslvserver" {
  vservername = citrixadc_lbvserver.tf_lbvserverssl.name
  ersa        = var.sslvsparam["ersa"]
  ssl3        = var.sslvsparam["ssl3"]
  tls1        = var.sslvsparam["tls1"]
  tls11       = var.sslvsparam["tls11"]
  depends_on = [citrixadc_lbvserver.tf_lbvserverssl]
} 