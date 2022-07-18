resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserverexample"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_intranetip6_binding" "tf_bind" {
  name        = citrixadc_vpnvserver.tf_vpnvserver.name
  intranetip6 = "2.3.4.5"
  numaddr     = "45"
}