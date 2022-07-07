resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_sharefileserver_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  sharefile = "3.3.4.3:90"
}