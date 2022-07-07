resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_staserver_binding" "tf_binding" {
  name           = citrixadc_vpnvserver.tf_vpnvserver.name
  staserver      = "http://www.example.com/"
  staaddresstype = "IPV4"
}