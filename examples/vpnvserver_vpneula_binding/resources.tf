resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpneula" "tf_vpneula" {
  name = "tf_vpneula"
}
resource "citrixadc_vpnvserver_vpneula_binding" "tf_bind" {
  name = citrixadc_vpnvserver.tf_vpnvserver.name
  eula = citrixadc_vpneula.tf_vpneula.name
}