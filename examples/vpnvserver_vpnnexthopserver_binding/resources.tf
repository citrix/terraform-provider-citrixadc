resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_exampleserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
  name        = "tf_vpnnexthopserver"
  nexthopip   = "2.2.1.5"
  nexthopport = "200"
}
resource "citrixadc_vpnvserver_vpnnexthopserver_binding" "tf_bind" {
  name          = citrixadc_vpnvserver.tf_vpnvserver.name
  nexthopserver = citrixadc_vpnnexthopserver.tf_vpnnexthopserver.name
}