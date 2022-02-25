resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name           = "tf_exampleserver"
  servicetype    = "SSL"
  ipv46          = "3.3.3.3"
  port           = 443
}
resource "citrixadc_vpnportaltheme" "tf_vpnportaltheme" {
  name      = "tf_vpnportaltheme"
  basetheme = "X1"
}
resource "citrixadc_vpnvserver_vpnportaltheme_binding" "tf_bind" {
  name        = citrixadc_vpnvserver.tf_vpnvserver.name
  portaltheme = citrixadc_vpnportaltheme.tf_vpnportaltheme.name
}