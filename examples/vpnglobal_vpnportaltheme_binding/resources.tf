resource "citrixadc_vpnportaltheme" "tf_vpnportaltheme" {
  name      = "tf_vpnportaltheme"
  basetheme = "X1"
}
resource "citrixadc_vpnglobal_vpnportaltheme_binding" "tf_bind" {
  portaltheme = citrixadc_vpnportaltheme.tf_vpnportaltheme.name
}