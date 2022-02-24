resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_newvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_appcontroller_binding" "tf_bind" {
  name          = citrixadc_vpnvserver.tf_vpnvserver.name
  appcontroller = "http://www.example.com"
}
