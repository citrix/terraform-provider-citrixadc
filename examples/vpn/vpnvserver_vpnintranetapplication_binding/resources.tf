resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
  intranetapplication = "tf_vpnintranetapplication"
  protocol            = "UDP"
  destip              = "2.3.6.5"
  interception        = "TRANSPARENT"
}
resource "citrixadc_vpnvserver_vpnintranetapplication_binding" "tf_bind" {
  name                = citrixadc_vpnvserver.tf_vpnvserver.name
  intranetapplication = citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
}
