resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_vpnportaltheme" "tf_vpnportaltheme" {
  name      = "tf_vpnportaltheme"
  basetheme = "X1"
}
resource "citrixadc_authenticationvserver_vpnportaltheme_binding" "tf_bind" {
  name = citrixadc_authenticationvserver.tf_authenticationvserver.name
  portaltheme = citrixadc_vpnportaltheme.tf_vpnportaltheme.name
}